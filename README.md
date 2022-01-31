# stock-high-low

Tiny go package for fetching high and low value of stock for any given period range using [Kite connect Historical
data APIs](https://kite.trade/docs/connect/v3/historical/).

## Installation
```
go get -u github.com/ranjanrak/stock-high-low
```

## Usage
```go
package main

import (
	contracthighlow "github.com/ranjanrak/contracthighlow"
)

func main() {
    // fetch 52 weeks high and low data
    result := contracthighlow.GetHighLow(contracthighlow.UserParam{
		ApiKey:      "your_api_key",
		AccessToken: "your_access_token",
		Token:       12942338,
		Day:         0,
		Month:       0,
		Year:        1,
    })
    
    fmt.Printf("%+v\n", result)
}
```

## Response
```
{High:4577.7 HighDate:2021-11-22 00:00:00 +0530 IST 
Low:3981.45 LowDate:2021-06-25 00:00:00 +0530 IST}
```
