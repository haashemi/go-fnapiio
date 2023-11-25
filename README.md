# Golang fortniteapi.io wrapper

A Simple Golang API wrapper for [FortniteAPI.io](https://fortniteapi.io).

Data models are written from the [official swagger documentation](https://fortniteapi.io/docs/).

This package is currently work-in-progress and there's no ETA for the stable release.

## Installation

```bash
$ go get -u github.com/haashemi/go-fnapiio
```

## Usage

```go
package main

import (
    "github.com/haashemi/go-fnapiio"
)

func main() {
    // initialize a new client with the token
    // from https://dashboard.fortniteapi.io
    client := fnapiio.New("my-api-token")
    // you can also pass your own http client
    // using fnapiio.NewWithHTTPClient

    // set a default language for the rest of the
    // requests, so you don't have to pass it every
    // time on each request. but you'll be able to
    // override it anytime on each request.
    client.SetLanguage(fnapiio.LangDE)

    // fetch the itemshop
    shop, err := client.Shop(fnapiio.ShopOpts{ Lang: "fr" })

    // ...the rest of the code...
}
```

## Contribution

All type of contributions on this package would be appreciated.
