package channel

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type StockPrice struct {
	Symbol string
	Price  float64
	Time   time.Time
}

func DemoChannel() {
	priceCh := make(chan StockPrice, 100)

	go priceProcessor(priceCh)

	go fetchStockPrices(priceCh, "APPL")
	go fetchStockPrices(priceCh, "GOOGL")
	go fetchStockPrices(priceCh, "MSFT")

	// Simulate some delay to allow fetching and processing
	time.Sleep(10 * time.Second)
}

// this channel is to send-only
func priceProcessor(priceCh <-chan StockPrice) {
	for price := range priceCh {
		fmt.Printf("Processing stock price : %s = %.2f at %s \n", price.Symbol, price.Price, price.Time)
		// Simulate processing time
		time.Sleep(500 * time.Millisecond)
	}
}

// fetchStockPrices will be called by a scheduler

// this channel to receiver-only
func fetchStockPrices(priceCh chan<- StockPrice, symbol string) {
	for {
		price := StockPrice{
			Symbol: symbol,
			Price:  rand.Float64() * 1000,
			Time:   time.Now(),
		}
		priceCh <- price
		// simulate delay between price updates
		time.Sleep(time.Second)
	}
}
