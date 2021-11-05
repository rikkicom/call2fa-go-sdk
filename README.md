# Go SDK for Call2FA

This is a library you can use for Rikkicom's service named as Call2FA 
(a call as the second factor in an authorization pipeline).

## Installation

Just install as the following:

```
go get github.com/rikkicom/call2fa-go-sdk
```

## Example

This simple code makes a new call to the +380631010121 number:

```go
package main

import (
	"fmt"
	"os"

	call2faSDK "github.com/rikkicom/call2fa-go-sdk"
)

func main() {
	// If you like, enable debug of HTTP requests, 0 to disable
	_ = os.Setenv("GOREQUEST_DEBUG", "1")

	// Configure the client
	cfg := call2faSDK.Config{
		Login:    "****",
		Password: "****",
	}

	// Create the Call2FA client
	c, err := call2faSDK.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// Configure variables
	phoneNumber := "+380631010121"
	callbackURL := "https://httpbin.org/post"

	// Do the request to start the call
	response, err := c.Call(phoneNumber, callbackURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Call ID:", response.CallID)
}
```

More examples are in the `examples` folder.

- Documentation: https://api.rikkicom.io/docs/en/call2fa/
- Documentation (in Ukrainian): https://api.rikkicom.io/docs/uk/call2fa/
- Documentation (in Russian): https://api.rikkicom.io/docs/ru/call2fa/