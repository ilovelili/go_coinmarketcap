# coinmarketcap
[Coinmarketcap](https://www.coinmarketcap.com/api) API Wrapper

[![Build Status](https://travis-ci.org/ilovelili/coinmarketcap.svg?branch=master)](https://travis-ci.org/ilovelili/coinmarketcap)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/ilovelili/coinmarketcap/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/ilovelili/coinmarketcap

# Installation

```
$ go get github.com/ilovelili/coinmarketcap
```

# Usage Example
```go
package main

import (
    "fmt"
    "github.com/ilovelili/coinmarketcap"
)

const (
    accessKey       = "hoge"
    secretAccessKey = "huga"
)

func main() {
    // accessKey and secretAccessKey is not necessary right now since coinmarketcap only provides public APIs at the moment
    client, _ := gocoinmarketcap.NewClient(accessKey, secretAccessKey)
    ctx := context.Background()
    tickers, _ := client.GetTickers(ctx)

    fmt.Printf("Tikcer = %+v\n", tikcer)
}
```

For detail, please check ``sample/cmd/coinmarketcap`` directory or [GoDoc](http://godoc.org/github.com/ilovelili/coinmarketcap).

# License
MIT
# Author
Min Ju <route666@live.cn>