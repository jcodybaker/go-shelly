# go-shelly
go-shelly is an unofficial native go client for the [Shelly Gen2 API](https://shelly-api-docs.shelly.cloud/gen2/). It compliments the [MongooseOS mgrpc library](https://pkg.go.dev/github.com/mongoose-os/mos/common/mgrpc).

## Maturity
This library is currently in active development (as of December 2023). It has meaningful gaps in testing and functionality. At this stage there is no guarantee of backwards compatibility. Once the project reaches a stable state, I will begin crafting releases with semantic versioning. 

## Usage
In most cases, using the library is as simple building an RPC channel and calling .Do().  See the [go doc](https://pkg.go.dev/github.com/jcodybaker/go-shelly) and [Shelly Gen2 API documentation](https://shelly-api-docs.shelly.cloud/gen2/).
```
import (
	"context"
	"fmt"
    "log"

	shelly "github.com/jcodybaker/go-shelly"
	"github.com/mongoose-os/mos/common/mgrpc"
)

const (
    rpcAddr = "http://192.168.1.20/rpc"
)

func main() {
    ctx := context.Background()
	c, err := mgrpc.New(ctx, rpcAddr, mgrpc.UseHTTPPost())
	if err != nil {
		log.Fatalf("establishing rpc channel: %v", err)
	}
	defer c.Disconnect(ctx)

    req := &shelly.ShellyGetStatusRequest{}
    statusResp, _, err := req.Do(ctx, c)
    if err != nil {
        log.Fatalf("querying device status: %v", err)
    }
    for i, sw := range statusResp.Switches {
        fmt.Printf("Switch %d status: %v\n", i, *sw.Output)
    }
}
```

## TODO
* More rigorous integration testing. Currently I have a Shelly Pro 4PM, Shelly Pro 3, Shelly Plug US, and Shelly Plus HT. All are controlling live workloads and thus I've been reluctant to test mutating actions outside the needs of my own projects.
* MQTT / WebSocket examples.
* mDNS / BLE integrations for discovery and bootstrapping.  Open question: Should these be in separate projects to avoid dependency bloat?
* Stabilize API
    * It's not always clear which Go types should be used for JSON Numeric values. More validation is needed to ensure these types are correctly defined.
    * In contrast to Golang, the Shelly API appears differentiate between no-value and `null`. For example, in some cases (ex. Wifi.SetConfig), a JSON key with `null` value will clear the existing configuration, while omitting the key/value entirely leaves it untouched.
* Complete API
    * Known missing: Scripts, Outbound WebSockets, Lights, ModBus, Covers, Voltmeters, Humidity, Temperature, DevicePower, Smoke, EM(data), EM1(data), PM1, UI.
    * Wrappers for easy software updates, and configuration of multi-line scripts and certificates.
* CI testing / dependabot / release pipeline.

## Contributing
Pull-requests and [issues](https://github.com/jcodybaker/go-shelly/issues) are welcome. Code should be formatted with gofmt, pass existing tests, and ideally add new testing. Test should include samples from live device request/response flows when possible.

## Legal

### Intellectual Property
This library and its authors (Cody Baker - cody@codybaker.com) have no affiliation with [Allterco Robotics](https://allterco.com/) (the maker of [Shelly](https://www.shelly.com/) devices) or [Cesanta Software Limited](https://cesanta.com/) (the maker of [MongooseOS](https://mongoose-os.com/)).  All trademarks are property of their respective owners.  Importantly, the "[Shelly](https://www.shelly.com/)" name and names of devices are trademarks of [Allterco Robotics](https://allterco.com/). [MongooseOS](https://mongoose-os.com/) and related trademarks are property of [Cesanta Software Limited](https://cesanta.com/).

### License 
Copyright 2023 [J Cody Baker](cody@codybaker.com) and licensed under the [MIT license](LICENSE.md).