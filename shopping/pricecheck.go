package shopping

import "tutorial.ground/shopping/db"

// PriceCheck is a public method
func PriceCheck(itemID int) (float64, bool) {
	item := db.LoadItem(itemID)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}

// priceCheck is a private method
func priceCheck(itemID int) (float64, bool) {
	item := db.LoadItem(itemID)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}

/*

Method ismi büyük harf ile yazılır ise (public)
küçük harf ile yazılır ise (private)

*/
