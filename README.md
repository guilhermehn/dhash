# dhash
> dHash algorithm in Go

## Install
~~~
go get github.com/guilhermehn/dhash
~~~

## Usage
~~~ go
package main

import (
  "fmt"
  "github.com/guilhermehn/dhash"
)

func main() {
  hash, err := dhash.Dhash("path/to/image", 8)

  if err != nil {
    fmt.Println("ERROR: %s", err.Error())
    return
  }

  fmt.Println(hash)
}
~~~

## Dependencies
+ [github.com/disintegration/imaging](//github.com/disintegration/imaging)

## License
MIT
