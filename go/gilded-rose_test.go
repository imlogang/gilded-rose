package main

import (
	"gildedRose/items"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func Test_GildedRose(t *testing.T) {
//	main()
//}

func Test_Item_Managment(t *testing.T) {
	testCases := []struct {
		testName string

		wantedAge     int
		wantedSellIn  int
		wantedQuality int
		testItem      Item
	}{
		// {
		// 	testName:      "Conjured Item Happy Path",
		// 	testItem:      &ConjuredItem{age: 0, sellIn: 10, quality: 2},
		// 	wantedAge:     1,
		// 	wantedSellIn:  9,
		// 	wantedQuality: 0,
		// },
		{
			testName:      "Conjured Item Happy Path",
			testItem:      items.MakeConjuredItem(age: 0, sellIn: 10, quality: 2, degredationRate: 2, ageRate: 1),
			wantedAge:     1,
			wantedSellIn:  10,
			wantedQuality: 0,
		},
		{
			testName:      "Legendary Item Happy Path",
			testItem:      &items.GenericItem{age: 0, sellIn: 10, quality: 2, degredationRate: 0, ageRate: 1},
			wantedAge:     1,
			wantedSellIn:  10,
			wantedQuality: 2,
		},
		// {
		// 	testName:      "Conjured Item Item Quality Does Not Go Below 0",
		// 	testItem:      &ConjuredItem{age: 1, sellIn: 10, quality: 1},
		// 	wantedAge:     2,
		// 	wantedSellIn:  9,
		// 	wantedQuality: 0,
		// },
		// {
		// 	testName:      "Standard Item Happy Path",
		// 	testItem:      &StandardItem{age: 0, sellIn: 10, quality: 1},
		// 	wantedAge:     1,
		// 	wantedSellIn:  9,
		// 	wantedQuality: 0,
		// },
		// {
		// 	testName:      "Standard Item Item Quality Does Not Go Below 0",
		// 	testItem:      &StandardItem{age: 0, sellIn: 10, quality: 0},
		// 	wantedAge:     1,
		// 	wantedSellIn:  9,
		// 	wantedQuality: 0,
		// },
		// {
		// 	testName:      "Standard Item Quality is never more than 50",
		// 	testItem:      &StandardItem{age: 0, sellIn: 10, quality: 50},
		// 	wantedAge:     1,
		// 	wantedSellIn:  9,
		// 	wantedQuality: 49,
		// },
		// {
		// 	testName:      "Lengdary Item Happy Path",
		// 	testItem:      &LegendaryItem{age: 0, sellIn: 10, quality: 50},
		// 	wantedAge:     1,
		// 	wantedSellIn:  10,
		// 	wantedQuality: 50,
		// },
		// {
		// 	testName:      "Standard Item Degredation twice after sell by date",
		// 	testItem:      &StandardItem{age: 0, sellIn: -1, quality: 50},
		// 	wantedAge:     1,
		// 	wantedSellIn:  -2,
		// 	wantedQuality: 48,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(tt *testing.T) {
			ItemManager(tc.testItem)
			assert.Equal(tt, tc.wantedAge, tc.testItem.Age())
			assert.Equal(tt, tc.wantedSellIn, tc.testItem.SellIn())
			assert.Equal(tt, tc.wantedQuality, tc.testItem.Quality())
		})
	}
}
