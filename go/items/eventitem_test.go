package items

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_EventItem_Management(t *testing.T) {
	testCases := []struct {
		testName string

		wantedAge     int
		wantedSellIn  int
		wantedQuality int
		testItem      Item
	}{
		{
			testName:      "Backstage Item Happy Path",
			testItem:      MakeEventItem("Backstage Item", 20, 0),
			wantedAge:     1,
			wantedSellIn:  19,
			wantedQuality: 1,
		},
		{
			testName:      "Backstage Item Happy Path Sell in 9 days",
			testItem:      MakeEventItem("Backstage Item", 9, 0),
			wantedAge:     1,
			wantedSellIn:  8,
			wantedQuality: 2,
		},
		{
			testName:      "Backstage Item Happy Path Sell in 4 days",
			testItem:      MakeEventItem("Backstage Item", 4, 0),
			wantedAge:     1,
			wantedSellIn:  3,
			wantedQuality: 3,
		},
		{
			testName:      "Backstage Item Happy Path Sell in 0 days",
			testItem:      MakeEventItem("Backstage Item", 0, 10),
			wantedAge:     1,
			wantedSellIn:  0,
			wantedQuality: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(tt *testing.T) {
			ItemManager(tc.testItem)
			assert.Equal(tt, tc.wantedAge, tc.testItem.Age(), fmt.Sprintf("Wanted age did not match: %#v", tc.testItem))
			assert.Equal(tt, tc.wantedSellIn, tc.testItem.SellIn(), fmt.Sprintf("Wanted SellIn did not match: %#v", tc.testItem))
			assert.Equal(tt, tc.wantedQuality, tc.testItem.Quality(), fmt.Sprintf("Wanted Quality did not match: %#v", tc.testItem))
		})
	}
}
