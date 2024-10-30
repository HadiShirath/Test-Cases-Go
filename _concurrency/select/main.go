package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// interruptionCH
	// sigCH

	sigCH := make(chan os.Signal, 1)
	signal.Notify(sigCH, syscall.SIGTERM, syscall.SIGINT) // terminate and interupt

	interruptionCH := make(chan bool)

	go func() {
		for {
			select {
			case <-interruptionCH:
				fmt.Printf("Task Goroutine diberhentikan...")
				return
			default:
				fmt.Println("Menjalankan Tugas...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	interruptionSignal := <-sigCH
	fmt.Printf("Interruption Signal Triggered: %v\n", interruptionSignal)
	fmt.Println("Shutting down the program gracefully...")
	close(interruptionCH) // Signal the goroutine to stop

	time.Sleep(time.Second) // Give goroutine time to finish
}
