package jobs

import (
	"daft-stats/db"
	"daft-stats/graph/model"
	"daft-stats/services"
	"math"
	"time"
)

func GenerateStats(properties []model.Property) {
	db := db.Connect()

	price_average := calculateAveragePrice(properties)
	price_high := calculateHighestPrice(properties)
	price_low := calculateLowestPrice(properties)
	price_distribution := calculatePriceDistribution(properties)

	// make a new stats object
	stats := model.Stat{
		Timestamp:         time.Now().String(),
		PriceAverage:      price_average,
		PriceHigh:         price_high,
		PriceLow:          price_low,
		PriceDistribution: price_distribution,
	}
	services.InsertStats(db, stats)

}

func calculateAveragePrice(properties []model.Property) float64 {
	var total_price float64
	for _, v := range properties {
		total_price += float64(v.Price)
	}
	return total_price / float64(len(properties))
}

func calculateHighestPrice(properties []model.Property) float64 {
	var highest_price float64
	for _, v := range properties {
		if float64(v.Price) > highest_price {
			highest_price = float64(v.Price)
		}
	}
	return highest_price
}

func calculateLowestPrice(properties []model.Property) float64 {
	lowest_price := 100000000.000
	for _, v := range properties {
		if float64(v.Price) < lowest_price && float64(v.Price) > 100 {
			lowest_price = float64(v.Price)
		}
	}
	return lowest_price
}

func calculatePriceDistribution(properties []model.Property) []int {
	distribution := []int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0}
	for _, v := range properties {
		index := int(math.Floor(float64(v.Price) / 500))
		if index <= 9 {
			distribution[index] += 1
		} else {
			distribution[10] += 1
		}
	}
	return distribution
}
