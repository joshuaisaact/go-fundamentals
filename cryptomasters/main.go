package main

import (
	"fmt"
	"sync"

	"frontendmasters.com/go/crypto/api"
)

func main() {
	// Array of currencies we want
	currencies := []string{"BTC", "ETH", "BCH"}

	// Defining a waitgroup
	var wg sync.WaitGroup

	// Looping through the array and creating a new goroutine per currency
	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			getCurrencyData(currency)
			wg.Done()
		}()
	}

	wg.Wait()

}

// New goroutine
func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
