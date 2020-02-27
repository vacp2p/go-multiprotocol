# go-multiprotocol

A lot of this code has been adapted from [multiaddr](https://github.com/multiformats/go-multiaddr).

Multiprotocol can be initialized using a CSV file, for the CSV format see the [specification](https://github.com/vacp2p/multiprotocol).

```go
err := Init("testdata/multiprotocol.csv")
if err != nil {
    print(err)
}
```

Protocols can also be added programatically using the ```AddProtocol``` function.