package main

import (
	"fmt"
	"gildedRose/items"
)


func main() {
	fmt.Println("Hello world!")
}

func ItemManager(item items.Item) {
	item.AgeItem()
}
