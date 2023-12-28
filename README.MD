# go-shelly
go-shelly is an unofficial native go client for the [Shelly Gen2 API](https://shelly-api-docs.shelly.cloud/gen2/). It compliments the [MongooseOS mgrpc library](https://pkg.go.dev/github.com/mongoose-os/mos/common/mgrpc).

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

## Legal

### Intellectual Property
This library and its authors (Cody Baker - cody@codybaker.com) have no affiliation with [Allterco Robotics](https://allterco.com/) (the maker of [Shelly](https://www.shelly.com/) devices) or [Cesanta Software Limited](https://cesanta.com/) (the maker of [MongooseOS](https://mongoose-os.com/)).  All trademarks are property of their respective owners.  Importantly, the "[Shelly](https://www.shelly.com/)" name and names of devices are trademarks of [Allterco Robotics](https://allterco.com/). [MongooseOS](https://mongoose-os.com/) and related trademarks are property of [Cesanta Software Limited](https://cesanta.com/).

### License 
Copyright 2023 [J Cody Baker](cody@codybaker.com) and licensed under the [MIT license](LICENSE.md).