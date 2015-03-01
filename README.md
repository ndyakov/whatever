# whatever

Package __whatever__ is a Go package that exports type Params with some useful methods on top of map[string]interface{}

[![BuildStatus](https://travis-ci.org/ndyakov/whatever.png)](https://travis-ci.org/ndyakov/whatever)
[![GoDoc](https://godoc.org/github.com/ndyakov/whatever?status.png)](https://godoc.org/github.com/ndyakov/whatever)
[![status](https://sourcegraph.com/api/repos/github.com/ndyakov/whatever/.badges/status.svg)](https://sourcegraph.com/github.com/ndyakov/whatever)

## Install

### Get the package
`go get github.com/ndyakov/whatever`

### Import in your source
`import "github.com/ndyakov/whatever"`

## Introduction

The main idea behind __whatever__ was for the type Params to be
used when unmarshaling JSON data from an `application/json` request,
but after adding the `Add` and `Remove` methods I think that this Params
type can be useful in variety of other cases as well.

There is a method that can transform the Params structure to
url.Values structure with specified prefix and suffix, for the
result can be used with Gorilla`s schema or Goji`s params packages.
Although some of the getters are useful for unmarshaled JSON
date you can also Add your own values to the Params structure.
You can also access nested Params objects.
If you need you can validate the existence of a specific key by
using the Required method.

### [type Params](http://godoc.org/github.com/ndyakov/whatever#Params)
_This may be outdated, please check the godoc for up-to-date documentation_
* [func NewFromJSON(jsonBody []byte) (Params, error)](http://godoc.org/github.com/ndyakov/whatever#NewFromJSON)
* [func (p Params) Add(key string, value interface{}) bool](http://godoc.org/github.com/ndyakov/whatever#Params.Add)
* [func (p Params) Empty() bool](http://godoc.org/github.com/ndyakov/whatever#Params.Empty)
* [func (p Params) Get(key string) string](http://godoc.org/github.com/ndyakov/whatever#Params.Get)
* [func (p Params) GetFloat(key string) float32](http://godoc.org/github.com/ndyakov/whatever#Params.GetFloat)
* [func (p Params) GetFloat32(key string) float32](http://godoc.org/github.com/ndyakov/whatever#Params.GetFloat32)
* [func (p Params) GetFloat64(key string) float64](http://godoc.org/github.com/ndyakov/whatever#Params.GetFloat64)
* [func (p Params) GetI(key string) interface{}](http://godoc.org/github.com/ndyakov/whatever#Params.GetI)
* [func (p Params) GetInt(key string) int](http://godoc.org/github.com/ndyakov/whatever#Params.GetInt)
* [func (p Params) GetInt64(key string) int64](http://godoc.org/github.com/ndyakov/whatever#Params.GetInt64)
* [func (p Params) GetInt8(key string) int8](http://godoc.org/github.com/ndyakov/whatever#Params.GetInt8)
* [func (p Params) GetP(key string) Params](http://godoc.org/github.com/ndyakov/whatever#Params.GetP)
* [func (p Params) GetSlice(key string) []interface{}](http://godoc.org/github.com/ndyakov/whatever#Params.GetSlice)
* [func (p Params) GetSliceInts(key string) []int](http://godoc.org/github.com/ndyakov/whatever#Params.GetSliceInts)
* [func (p Params) GetSliceStrings(key string) []string](http://godoc.org/github.com/ndyakov/whatever#Params.GetSliceStrings)
* [func (p Params) GetString(key string) string](http://godoc.org/github.com/ndyakov/whatever#Params.GetString)
* [func (p Params) GetTime(key string) time.Time](http://godoc.org/github.com/ndyakov/whatever#Params.GetTime)
* [func (p Params) Remove(key string)](http://godoc.org/github.com/ndyakov/whatever#Params.Remove)
* [func (p Params) Required(keys ...string) error](http://godoc.org/github.com/ndyakov/whatever#Params.Required)
* [func (p Params) URLValues(prefix, suffix string) url.Values](http://godoc.org/github.com/ndyakov/whatever#Params.URLValues)

## Example

### NewFromJSON

```go

import (
  "fmt"

  "github.com/ndyakov/whatever"
)

var body = []byte(`
{
  "int": -10,
  "string": "test",
  "time": "2015-02-20T21:22:23.24Z",
  "nestedParams": {
    "arrayStrings": ["one","two","three"],
    "arrayInts": [1,2,3,4],
  }
}
`)

func main() {
  p := whatever.NewFromJSON(body)
  fmt.Println(p.GetInt("int")) // -10
  fmt.Println(p.GetP("nestedParams").GetSliceStrings("arrayStrings")[1]) // two
}
```

### Params{}

```go
func sum(p whatever.Params) int {
  x := p.GetInt("x")
  // or x := p.GetI("x").(int)
  y := p.GetInt("y")
  // or y := p.GetI("y").(int)
  return x + y
}

func main() {
  p := Params{x: 10, y: 5}
  fmt.Println(sum(p)) // 15
}
```

## Contributions

Before contributing please execute:
* gofmt
* golint
* govet
