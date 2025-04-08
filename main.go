package main

import "github.com/Ceilovee/Bootcamp-go/internal"

func main() {
	p, err := internal.LoadSliceProducts()
	if err != nil {
		panic(err)
	}
	internal.RaiseServer(&p)
}
