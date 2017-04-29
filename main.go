package main

import (
	"fmt"
	"compress/compression"

	"gopkg.in/mgo.v2/bson"
)


const DATABASE = "compression"
const COLLECTION = "compression"

func main() {
	driver, err := compression.NewDriver(DATABASE, COLLECTION)
	if err != nil {
		fmt.Println("Something went wrong.")
		fmt.Println(nil)
	}

	driver.AddString()
	findDict := bson.M{
		"_id": "compression",
	}
	fmt.Println(driver.GetAttribute(findDict))
}