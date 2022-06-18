package models

type Property struct {
	Price    int    `bson: price`
	Location string `bson: location`
	URL      string `bson: url`
	DaftId   int    `bson: id`
	Entered  string `bson: entered`
	Views    int    `bson: views`
	Type     string `bson: type`
	Bathroom int    `bson: bathroom`
	Bed      int    `bson: bed`
}
