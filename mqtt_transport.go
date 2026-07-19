package shelly

// This file adds an MQTT transport for the mgrpc RPC layer that shares one
// already-connected mqtt.Client across many devices on the same broker,
// rather than dialing a new MQTT connection per device (which is what
// mongoose-os/mos's own codec.MQTT does). This matches how Shelly devices
// are actually deployed: many devices, one broker.
//
// Ported from github.com/jcodybaker/shellyctl's pkg/discovery/mqtt_codec.go
// (private to that package), which is itself adapted from
// github.com/mongoose-os/mos/common/mgrpc/codec's MQTT codec.
//
// Original License:
//
// Copyright (c) 2014-2019 Cesanta Software Limited
// All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mongoose-os/mos/common/mgrpc"
	"github.com/mongoose-os/mos/common/mgrpc/codec"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

// mqttCodec implements mgrpc/codec.Codec for RPC calls to one Shelly device
// (dst, e.g. "shellypro4pm-c8f09e87d088"), reusing a shared mqtt.Client.
// Requests publish to "<dst>/rpc"; responses are read from
// "<dst>/rpc-resp/<mqtt-client-id>/rpc".
type mqttCodec struct {
	dst         string
	closeNotify chan struct{}
	rchan       chan frame.Frame
	cli         mqtt.Client
	closeOnce   sync.Once
	isTLS       bool
	pubTopic    string
	subTopic    string
	mu          sync.Mutex
}

// NewMQTTCodec returns an mgrpc codec.Codec for RPC calls to the Shelly
// device identified by dst (its Shelly ID, e.g.
// "shellypro4pm-c8f09e87d088"), over an already-connected, shared
// mqtt.Client. Callers wanting go-shelly's typed request/response helpers
// (e.g. SwitchSetRequest.Do) should wrap the result with mgrpc.Serve or
// pass it to NewMQTTTransport.
func NewMQTTCodec(dst string, cli mqtt.Client) (codec.Codec, error) {
	co := cli.OptionsReader()
	subTopic := fmt.Sprintf("%s/rpc-resp/%s", dst, co.ClientID())
	c := &mqttCodec{
		dst:         dst,
		closeNotify: make(chan struct{}),
		rchan:       make(chan frame.Frame),
		subTopic:    subTopic,
		pubTopic:    fmt.Sprintf("%s/rpc", dst),
		isTLS:       len(co.Servers()) > 0 && co.Servers()[0].Scheme == "tcps",
		cli:         cli,
	}
	token := cli.Subscribe(subTopic+"/rpc", 1 /* qos */, c.onMessage)
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, fmt.Errorf("subscribing to mqtt %q: %w", subTopic+"/rpc", err)
	}
	return c, nil
}

// NewMQTTTransport wraps NewMQTTCodec with mgrpc.Serve so go-shelly's typed
// request Do() methods work transparently over MQTT against dst, using an
// already-connected, shared mqtt.Client. ctx bounds the codec's receive
// loop; canceling it (or calling the returned mgrpc.MgRPC's Disconnect)
// unsubscribes and stops delivering responses/notifications.
func NewMQTTTransport(ctx context.Context, dst string, cli mqtt.Client) (mgrpc.MgRPC, error) {
	c, err := NewMQTTCodec(dst, cli)
	if err != nil {
		return nil, err
	}
	return mgrpc.Serve(ctx, c), nil
}

func (c *mqttCodec) onMessage(_ mqtt.Client, msg mqtt.Message) {
	f := &frame.Frame{}
	if err := json.Unmarshal(msg.Payload(), f); err != nil {
		return // malformed payload; nothing sensible to do but drop it.
	}
	select {
	case c.rchan <- *f:
	case <-c.closeNotify:
	}
}

func (c *mqttCodec) Close() {
	c.closeOnce.Do(func() {
		token := c.cli.Unsubscribe(c.subTopic + "/rpc")
		token.Wait()
		close(c.closeNotify)
	})
}

func (c *mqttCodec) CloseNotify() <-chan struct{} { return c.closeNotify }

func (c *mqttCodec) String() string { return fmt.Sprintf("[mqttCodec to %s]", c.dst) }

func (c *mqttCodec) Info() codec.ConnectionInfo {
	return codec.ConnectionInfo{
		IsConnected: c.cli.IsConnected(),
		TLS:         c.isTLS,
		RemoteAddr:  c.dst,
	}
}

func (c *mqttCodec) MaxNumFrames() int { return -1 }

func (c *mqttCodec) Recv(ctx context.Context) (*frame.Frame, error) {
	select {
	case f := <-c.rchan:
		return &f, nil
	case <-c.closeNotify:
		return nil, io.EOF
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (c *mqttCodec) Send(ctx context.Context, f *frame.Frame) error {
	f.Src = c.subTopic
	f.Dst = c.dst
	msg, err := json.Marshal(f)
	if err != nil {
		return fmt.Errorf("marshalling mqtt rpc frame: %w", err)
	}
	token := c.cli.Publish(c.pubTopic, 1 /* qos */, false /* retained */, msg)
	token.Wait()
	if err := token.Error(); err != nil {
		return fmt.Errorf("mqtt publish error: %w", err)
	}
	return nil
}

func (c *mqttCodec) SetOptions(opts *codec.Options) error {
	return fmt.Errorf("mqttCodec: SetOptions not implemented")
}
