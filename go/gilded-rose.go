package main

import (
	"gildedRose/items"
)

//type Item struct {
//	name            string
//	sellIn, quality int
//}

// var itemsSet = []items.Item{
// 	&items.GenericItem{
// 		name:    "+5 Dexterity Vest",
// 		sellIn:  10,
// 		quality: 20,
// 		age:     10,
// 	},
// 	&items.GenericItem{
// 		name:    "Aged Brie",
// 		sellIn:  2,
// 		quality: 0,
// 		age:     10,
// 	},
// 	&items.GenericItem{
// 		name:    "Elixir of the Mongoose",
// 		sellIn:  5,
// 		quality: 7,
// 		age:     10,
// 	},
// 	&items.GenericItem{
// 		name:    "Sulfuras, Hand of Ragnaros",
// 		sellIn:  0,
// 		quality: 80,
// 		age:     10,
// 	},
// 	&items.GenericItem{
// 		name:    "Backstage passes to a TAFKAL80ETC concert",
// 		sellIn:  15,
// 		quality: 20,
// 		age:     10,
// 	},
// 	&items.GenericItem{
// 		name:    "Conjured Mana Cake",
// 		sellIn:  3,
// 		quality: 6,
// 		age:     10,
// 	},
// }

func main() {
	// for _, v := range items {
	// 	v.Display()
	// 	v.Age()
	// 	v.Display()
	// }

	// for _, n := range items {
	// 	fmt.Println)
	// }

	//	fmt.Println("OMGHAI!")
	//
	//fmt.Print(items)
	//
	// for _, g := range items {
	// 	fmt.Println("Starting values")
	// 	g.Display()
	// 	fmt.Println("New Values")
	// 	g.Display()
	// 	fmt.Println("----------------")
	// }
}

// func (s *items.GenericItem) GildedRose() {
// 	if s.Name() != "Aged Brie" && s.Name() != "Backstage passes to a TAFKAL80ETC concert" {
// 		if s.Quality() > 0 {
// 			if s.Name() != "Sulfuras, Hand of Ragnaros" && s.Name() != "Conjured Mana Cake" {
// 				s.quality -= 1
// 			}
// 			if s.Name() == "Conjured Mana Cake" {
// 				s.quality -= 2
// 			}
// 		}
// 	} else {
// 		if s.Quality() < 50 {
// 			s.quality += 1
// 			if s.Name() == "Backstage passes to a TAFKAL80ETC concert" {
// 				if s.SellIn() < 11 {
// 					if s.Quality() < 50 {
// 						s.quality += 1
// 					}
// 				}
// 				if s.SellIn() < 6 {
// 					if s.Quality() < 50 {
// 						s.quality += 1
// 					}
// 				}
// 			}
// 		}
// 	}

// 	if s.Name() != "Sulfuras, Hand of Ragnaros" {
// 		s.sellIn -= 1
// 	}

// 	if s.SellIn() < 0 {
// 		if s.Name() != "Aged Brie" {
// 			if s.Name() != "Backstage passes to a TAFKAL80ETC concert" {
// 				if s.Quality() > 0 {
// 					if s.Name() != "Sulfuras, Hand of Ragnaros" {
// 						s.quality -= 1
// 					}
// 				}
// 			} else {
// 				s.quality = s.quality - s.quality
// 			}
// 		} else {
// 			if s.Quality() < 50 {
// 				s.quality += 1
// 			}
// 		}
// 	}
// }

func ItemManager(item items.Item) {
	item.AgeItem()
}
