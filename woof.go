package shelly

import (
	"context"

	"github.com/mongoose-os/mos/common/mgrpc/codec"
	"github.com/mongoose-os/mos/common/mgrpc/frame"
)

func do() error {
	ctx := context.Background()
	dst := "hi"
	c, err := codec.MQTT(mqURL, nil, &codec.MQTTCodecOptions{})
	if err != nil {
		return err
	}
	f := frame.NewRequestFrame("a", "b", "blah", &frame.Command{}, false)
	c.Send(ctx, f)
}
