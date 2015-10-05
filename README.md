# dhash
> dHash algorithm in Go

## Install
~~~ bash
go get github.com/guilhermehn/dhash
~~~
Or with [gopkg](http://labix.org/gopkg.in):
~~~ bash
go get gopkg.in/guilhermehn/dhash.v1
~~~

## Usage
Two parameters are needed: the path to the image and
the size that the image will be resized to, that will reflect
into the hash size. **The recommended size is 8**.

~~~ go
Dhash("image.jpg", 8)
~~~

The image will be resized to 8x8 and the hash
size will be `size*2`: `"8899aabbccddeeff"`.

### Example:

~~~ go
package main

import (
  "fmt"
  "github.com/guilhermehn/dhash"
  // or if you installed the gopkg version
  // "gopkg.in/guilhermehn/dhash.v1"
)

func main() {
  pathToImage := "path/to/image"
  detailLevel := 8

  hash, err := dhash.Dhash(pathToImage, detailLevel)

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
