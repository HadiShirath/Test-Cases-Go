package channel

import (
	"fmt"
	"time"
)

func Demo() {
	fmt.Println("Program Dijalankan ...")
	// flow:
	// - buatkan channel
	// - buatkan simulasi penerimaan data lewat channel
	// - buatkan simulasi pengiriman data lewat channel
	// - tunggu sampai program selesai

	messageCh := make(chan string, 4)

	// receiver
	go func() {
		for {
			messageData := <-messageCh
			fmt.Printf("data diterima : %s \n", messageData)
		}
	}()

	// sender
	go func() {
		for i := 1; i <= 12; i++ {
			fmt.Printf("data ke : %d di kirim \n", i)
			messageCh <- fmt.Sprintf("data dari goroutine ke : %d", i)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Program Selesai")
}
