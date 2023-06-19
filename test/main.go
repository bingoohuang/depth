package main

import (
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`


func main() {
	fmt.Println(os.Getwd())
	value := gjson.Get(json, "name.last")
	println(value.String())
}