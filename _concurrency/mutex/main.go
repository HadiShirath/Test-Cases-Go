package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var stok int = 100

func beliProduk(seq int, jumlah int) {
	mutex.Lock()
	defer mutex.Unlock()

	if stok >= jumlah {
		stok -= jumlah
		fmt.Printf("Pembelian ke : %d berhasil, dengan jumlah : %d\n", seq, jumlah)
	} else {
		fmt.Printf("Pembelian ke : %d gagal, stok tidak cukup\n", seq)
	}

}

func main() {
	for i := 0; i < 10; i++ {
		go beliProduk(i, 20)
	}

	time.Sleep(3 * time.Second)
}
