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

	attribute := "compression"
	entry := &bson.M{
		"_id": attribute,
		"counter": 3,
		"values": bson.M{
			"Romania": 0,
			"Bucharest": 1,
			"Azimut": 2,
		},
	}
	driver.AddString(entry)
	fmt.Println(driver.GetAttribute(attribute))
}