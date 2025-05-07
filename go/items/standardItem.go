package items

import "fmt"

type Item interface {
	AgeItem()
	Display()
	Name() string
	Quality() int
	SellIn() int
	Age() int
}

const (
	STANDARD         = iota
	LEGENDARY        = iota
	CONJURED         = iota
	CHEESE           = iota
	BACKSTAGE_PASSES = iota
)

type GenericItem struct {
	name                 string
	sellIn, quality, age int
	degredationRate      int
	ageRate              int
	itemType             int
}

func (s *GenericItem) AgeItem() {
	s.age += s.ageRate
	s.sellIn -= s.ageRate
	if s.quality > 0 && s.quality < 50 {
		if s.age > s.sellIn {
			s.quality -= (s.degredationRate * 2)
		} else {
			s.quality -= s.degredationRate
		}
	}
}

func (s *GenericItem) Display() {
	fmt.Printf("Name: %s\nAge: %d\nSell in: %d\nQuality: %d\n\n", s.name, s.age, s.sellIn, s.quality)
}

func (s *GenericItem) Name() string {
	return s.name
}

func (s *GenericItem) Quality() int {
	return s.quality
}

func (s *GenericItem) SellIn() int {
	return s.sellIn
}

func (s *GenericItem) Age() int {
	return s.age
}

func MakeStandardItem(name string, sellIn, quality int) *GenericItem {
	return &GenericItem{
		name:            name,
		sellIn:          sellIn,
		quality:         quality,
		degredationRate: 1,
		ageRate:         1,
		itemType:        STANDARD,
	}
}

func MakeConjuredItem(name string, sellIn, quality int) *GenericItem {
	return &GenericItem{
		name:            name,
		sellIn:          sellIn,
		quality:         quality,
		degredationRate: 2,
		ageRate:         1,
		itemType:        CONJURED,
	}
}

func MakeLegendaryItem(name string, sellIn, quality int) *GenericItem {
	return &GenericItem{
		name:            name,
		sellIn:          sellIn,
		quality:         quality,
		degredationRate: 0,
		ageRate:         0,
		itemType:        LEGENDARY,
	}
}

func MakeCheeseItem(name string, sellIn, quality int) *GenericItem {
	return &GenericItem{
		name:            name,
		sellIn:          sellIn,
		quality:         quality,
		degredationRate: -1,
		ageRate:         1,
		itemType:        CHEESE,
	}
}

func MakeBACKSTAGE_Item(name string, sellIn, quality int) *GenericItem {
	return &GenericItem{
		name:            name,
		sellIn:          sellIn,
		quality:         quality,
		degredationRate: -1,
		ageRate:         1,
		itemType:        BACKSTAGE_PASSES,
	}
}
