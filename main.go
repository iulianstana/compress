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
		"Romania": 9,
		"Bucharest": 3,
		"Azimut": 121,
	}
	driver.AddString(entry)
	fmt.Println(driver.LoadAttribute(attribute))
	fmt.Println(driver.UpdateValue(attribute, 5))
	fmt.Println(driver.UpdateValue(attribute, 100))
	fmt.Println(driver.UpdateValue(attribute, -1))
}