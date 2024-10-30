package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printNumbers(1, &wg)

	wg.Add(1)
	go printNumbers(2, &wg)

	wg.Wait()
	fmt.Println("Program Selesai")
}

func printNumbers(jobID int, wg *sync.WaitGroup) {
	for i := 1; i < 5; i++ {
		fmt.Printf("Job ID : %d Menjalankan task : %d \n", jobID, i)
	}

	wg.Done()

}
