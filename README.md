# go-multiprotocol

![Version](https://img.shields.io/github/tag/vacp2p/go-multiprotocol.svg)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://godoc.org/github.com/vacp2p/go-multiprotocol) 
[![Go Report Card](https://goreportcard.com/badge/github.com/vacp2p/go-multiprotocol)](https://goreportcard.com/report/github.com/vacp2p/go-multiprotocol)
[![Build Status](https://travis-ci.com/vacp2p/go-multiprotocol.svg?branch=master)](https://travis-ci.com/vacp2p/go-multiprotocol)

A lot of this code has been adapted from [multiaddr](https://github.com/multiformats/go-multiaddr).

Multiprotocol can be initialized using a CSV file, for the CSV format see the [specification](https://github.com/vacp2p/multiprotocol).

```go
package main

import mp "github.com/vacp2p/go-multiprotocol"

err := mp.Init("testdata/multiprotocol.csv")
// err non-nil when parsing failed.

mp, err := mp.NewMultiprotocol("/vac/waku/2/store/2/relay/2")
// err non-nil when parsing failed.
```

Protocols can also be added programatically using the ```AddProtocol``` function.
