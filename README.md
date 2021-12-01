# vonage-go-sdk
unofficial sdk for vonage - communication platform - implemented by Go.

## Milestone
- 0.0.1: send verify code
- 0.0.2: check verify code
- 0.0.3: cancel verify operation
- future
  - other apis

## Setup
It's recommended to use `go modules`. Import this package on your project, then operate `go mod tidy`.

## Usage
### Send verify code
```go
import "github.com/tenkoh/vonage-go-sdk"

func main(){
    client, _ := vonage.NewClient(
        vonage.ApiKey("YOUR_API_KEY"),
        vonage.ApiSecret("YOUR_API_SECRET"),
    )
    resp, _ := client.GenerateVerifyClient().Verify(
		vonage.VerifyNumber("PHONE_NUMBER"),
		vonage.VerifyBrand("YOUR_BRAND_NAME"),
	)
}
```

You can omit `vonage.ApiKey` and `vonage.ApiSecret` if you have exported them as emvironment variables. (VONAGE_API_KEY and VONAGE_API_SECRET)

## License
MIT

note: the other vonage sdks are published under BSD-3 license.

Note: This package references official sdk implemented by Python3. The official sdk is published under Apache license.
