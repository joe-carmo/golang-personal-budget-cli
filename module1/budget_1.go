package module1

// Budget stores budget information
struct Budget {
	Max float32,
	Items []Item,
}
// Item stores item information
struct Item {
	Description string,
	Price float32,
}