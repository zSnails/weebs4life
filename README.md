# What is this?
This is just a <i>smol</i> api wrapper for [weebs4life](https://weebs4life.ga) an API to retrieve the most <i>weeb</i> shit you'd imagine (like srsly) 

# Installation
```bash
$ go get github.com/zSnails/weebs4life
```

# Examples

```go
package main

import (
    "fmt"
    "github.com/zSnails/weebs4life"
)

//Initialize the w e e b
var weebClient = weebs4life.Weeb{}

func main() {
    //get a hug
	hug, err := weebClient.Hug()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hug: %v\n", hug.URL)
    //get a kiss
	kiss, err := weebClient.Kiss()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Kiss: %v\n", kiss.URL)
    //get a neko
	neko, err := weebClient.Neko()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Neko: %v\n", neko.URL)
    //get slapped lmao
	slap, err := weebClient.Slap()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Slap: %v\n", slap.URL)
}

```