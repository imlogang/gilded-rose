package items

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Item_Management(t *testing.T) {
	testCases := []struct {
		testName string

		wantedAge     int
		wantedSellIn  int
		wantedQuality int
		testItem      Item
	}{
		{
			testName:      "Conjured Item Happy Path",
			testItem:      MakeConjuredItem("Conjured Item", 10, 2),
			wantedAge:     1,
			wantedSellIn:  9,
			wantedQuality: 0,
		},
		{
			testName:      "Legendary Item Happy Path",
			testItem:      MakeLegendaryItem("Legendary Item", 10, 2),
			wantedAge:     0,
			wantedSellIn:  10,
			wantedQuality: 2,
		},
		{
			testName:      "Standard Item Happy Path",
			testItem:      MakeStandardItem("Standard Item", 10, 1),
			wantedAge:     1,
			wantedSellIn:  9,
			wantedQuality: 0,
		},
		{
			testName:      "Cheese Item Happy Path",
			testItem:      MakeCheeseItem("Cheese Item", 10, 1),
			wantedAge:     1,
			wantedSellIn:  9,
			wantedQuality: 2,
		},
		{
			testName:      "Backstage Item Happy Path",
			testItem:      MakeBACKSTAGE_Item("Backstage Item", 20, 0),
			wantedAge:     1,
			wantedSellIn:  19,
			wantedQuality: 1,
		},
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
