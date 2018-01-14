config [![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)[![Coverage Status](https://coveralls.io/repos/github/klokare/config/badge.svg?branch=master)](https://coveralls.io/github/klokare/config?branch=master) [![GoDoc](https://godoc.org/github.com/klokare/config?status.svg)](https://godoc.org/github.com/klokare/config)
====

Helpers useful for configuring Go structs.

Original conceived to provide alternative ways for configuring experiments in [evo](https://github.com/klokare/evo), this package exists outside that library for two reasons:
1. To keep the evo library itself as independent as possible
2. To provide this functionality without requiring the full evo source

## Install
To start using the config library, install Go and run `go get`:

```sh
$ go get github.com/klokare/config/...
```

## Usage
Each helper provides its own New method(s). Some rely on an io.Reader for the source configuration, others might use environment variables, SQL databases, or key-value stores. Once instantiated, though, the use will be the same:

```go
// Create the new configurer
var cfg *json.Configurer
if cfg, err = json.NewFromFile("example.json"); err != nil {
    ...
}

// Configure other components
var baker   Baker
var butcher Butcher
var maker   candlestick.Maker
if err = cfg.Configure(&baker, &butcher, &maker); err != nil {
    ...
}
```
