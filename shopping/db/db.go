package db

// Item is
type Item struct {
	Price float64
}

// LoadItem is
func LoadItem(id int) *Item {
	return &Item{
		Price: 9.0001,
	}
}
