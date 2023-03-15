# botastic-go

A Go SDK package for botastic service. https://developers.pando.im/references/botastic/api.html

## Example

```go
package main

import (
	"context"

	"github.com/pandodao/botastic-go"
)

func main() {
	client := botastic.New("YOUR_APP_ID", "YOUR_APP_SECRET", botastic.WithHost("host..."), botastic.WithDebug(true))
	resp, err := client.SearchIndexes(context.Background(), botastic.SearchIndexesRequest{
		Keywords: "hello",
		N:        10,
	})
	if err != nil {
		panic(err)
	}

	for _, index := range resp.Items {
		println(index.ObjectID, index.Data)
	}
}
```

## Options

* WithHost - Set the custom host, the default host is our production environment service address.
* WithDebug - The relevant request and response logs will be printed out.
* WithLogger - Custom logger for debug log.

## License
[MIT License](https://github.com/pandodao/botastic-go/blob/main/LICENSE)

