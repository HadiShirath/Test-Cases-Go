package main

import (
	"sync"
)

func main() {
	// var wg sync.WaitGroup

	// wg.Add(3)

	// func1(&wg)
	// func2(&wg)
	// func3(&wg)
}

func fetchPricing() {
	// start
	// fetch Indosat API
	// fetch Telkomsel API
	// fetch XL API
	// done

	var wg sync.WaitGroup

	wg.Add(1)
	indosatPricingResult, indosatErr := fetchIndosatAPI(&wg)

	wg.Add(1)
	telkomselPricingResult, telkomselErr := fetchTelkomselAPI(&wg)

	wg.Add(1)
	xlPricingResult, xlErr := fetchXLAPI(&wg)

	// cache all pricing result into Redis
	_ = indosatPricingResult
	_ = telkomselPricingResult
	_ = xlPricingResult

	_ = indosatErr
	_ = telkomselErr
	_ = xlErr
}

func fetchIndosatAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	return
}

func fetchTelkomselAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	return
}

func fetchXLAPI(wg *sync.WaitGroup) (data struct{}, err error) {
	defer wg.Done()

	return
}

// func func1(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// run the logic here
// 	println("execute logic 1")
// }

// func func2(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// run the logic here
// 	println("execute logic 2")
// }

// func func3(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// run the logic here
// 	println("execute logic 3")
// }
