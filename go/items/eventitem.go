package items

import "fmt"

type EventItem struct {
	name                 string
	sellIn, quality, age int
	ageRate              int
	itemType             int
}

func (s *EventItem) Display() {
	fmt.Printf("Name: %s\nAge: %d\nSell in: %d\nQuality: %d\n\n", s.name, s.age, s.sellIn, s.quality)
}

func (s *EventItem) Name() string { return s.name }

func (s *EventItem) Quality() int {
	return s.quality
}

func (s *EventItem) SellIn() int {
	return s.sellIn
}

func (s *EventItem) Age() int {
	return s.age
}

func (s *EventItem) AgeItem() {
	//Increases in quality
	//Increases by 2 when there's less than 10 days
	//and by 3 when there are 5 days or less
	//Drops to 0 when sellIn is 0
	//Quality will never go below 0, and never greater than 50.
	s.age += s.ageRate
	if s.sellIn > 0 {
		s.sellIn -= s.ageRate
	}
	if s.quality >= 0 && s.quality <= 50 {
		switch {
		case s.sellIn <= 10 && s.sellIn > 5:
			s.quality += 2
		case s.sellIn <= 5 && s.sellIn > 0:
			s.quality += 3
		case s.sellIn == 0:
			s.quality = 0
		default:
			s.quality += 1
		}
	}
}

func MakeEventItem(name string, sellIn, quality int) *EventItem {
	return &EventItem{
		name:     name,
		sellIn:   sellIn,
		quality:  quality,
		ageRate:  1,
		itemType: BACKSTAGE_PASSES,
	}
}
