package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("program dimulai ...")
	// flow
	// 1. create bookstore
	BookStore := NewBookStore()

	// 2. Simulate multiple readers
	for i := 0; i < 5; i++ {
		go func() {
			book := BookStore.getBookDetail(1)
			if book != nil {
				fmt.Printf("reader ke : %d, buku dengan id : %d ditemukan, quantity : %d\n", i, 1, book.Quantity)
			} else {
				fmt.Printf("reader ke %d, buku tidak ditemukan \n", i)
			}
		}()
	}

	// 3. Simulate Write
	go func() {
		fmt.Println("writer : mengupdate jumlah buku...")
		BookStore.UpdateBookQuantity(1, -1)
		fmt.Println("writer : berhasil mengupdate jumlah buku")
		time.Sleep(200 * time.Millisecond)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("program selesai ...")
}

type Book struct {
	Title    string
	Quantity int
}

type BookStore struct {
	books   map[int]*Book
	rwmutex sync.RWMutex
}

func NewBookStore() *BookStore {
	return &BookStore{
		books: map[int]*Book{
			1: {
				Title:    "Go Programming",
				Quantity: 10,
			},
			2: {
				Title:    "Concurrency in Go",
				Quantity: 5,
			},
		},
	}
}

func (bs *BookStore) getBookDetail(id int) *Book {
	bs.rwmutex.RLock()
	defer bs.rwmutex.RUnlock()

	// di map bisa mengecek found
	book, found := bs.books[id]
	if !found {
		fmt.Printf("book with id : %d not found \n", id)
		return nil
	}

	return &Book{
		Title:    book.Title,
		Quantity: book.Quantity,
	}

}

func (bs *BookStore) UpdateBookQuantity(id int, change int) {
	bs.rwmutex.Lock()
	defer bs.rwmutex.Unlock()

	book, found := bs.books[id]
	if !found {
		fmt.Printf("book with id : %d not found \n", id)
		return
	}

	book.Quantity += change
}
