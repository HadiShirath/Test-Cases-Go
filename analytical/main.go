package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// PIPELINE PATTERN

	// simulate to fetching order
	// simulate to filtering order
	// simulate to analysing order item
	// simulate storing the analysis result to big query

	// fetch -> filter -> analyse -> store

	orderCH := make(chan Order)
	filteredOrderCH := make(chan Item)
	analyseReportCH := make(chan AnalysisReport)

	// run stage of pipeline concurrently
	go fetchOrder(orderCH)
	go filtered(orderCH, filteredOrderCH)
	go analyseOrder(filteredOrderCH, analyseReportCH)
	go storeAnalysisReport(analyseReportCH)

	select {}
}

type (
	Order struct {
		ID    int
		Items []Item
	}

	Item struct {
		ItemID       int
		Category     string // to be filtered, spesific for "digital" category
		ProviderName string // telkomsel, indosat, xl
		Price        float64
	}

	AnalysisReport struct {
		Category     string
		AveragePrice float64
		MinPrice     float64
		MaxPrice     float64
	}
)

func fetchOrder(orderCH chan<- Order) {
	for i := 0; ; i++ {
		orderData := Order{
			ID: i,
			Items: []Item{
				{
					ItemID:       i*10 + 2,
					Category:     "digital",
					ProviderName: "telkomsel",
					Price:        rand.Float64() * 10000,
				},
			},
		}

		time.Sleep(1 * time.Second)
		orderCH <- orderData
		fmt.Println("fetched the order")
	}

}

func filtered(orderCH <-chan Order, filteredOrderCH chan<- Item) {
	for order := range orderCH {
		for _, item := range order.Items {
			if item.Category == "digital" {
				time.Sleep(200 * time.Millisecond)
				filteredOrderCH <- item
				fmt.Println("filtered order detail")
			}
		}
	}
}

func analyseOrder(filteredOrderCH <-chan Item, analyseReportCH chan<- AnalysisReport) {
	for item := range filteredOrderCH {
		// do the analysis here
		result := AnalysisReport{
			Category:     item.Category,
			MinPrice:     item.Price,     // to simplify the calculation
			MaxPrice:     item.Price * 2, // to simplify the calculation
			AveragePrice: item.Price,     // to simplify the calculation
		}
		analyseReportCH <- result
		fmt.Println("analyzed the order detail")
	}
}

func storeAnalysisReport(analyseReportCH <-chan AnalysisReport) {
	for report := range analyseReportCH {
		// store the big query
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("stored the analysis report to BQ with result %+v\n\n", report)
	}

}
