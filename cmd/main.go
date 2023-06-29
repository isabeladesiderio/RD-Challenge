package main

import (
	"fmt"

	"github.com/ResultadosDigitais/developer_challenges/go/internal/csbalancing"
)

type Entity struct {
	ID    int
	Score int
}

func main() {
	// Example usage of CustomerSuccessBalancing function
	customerSuccess := []csbalancing.Entity{
		{ID: 1, Score: 500},
		{ID: 2, Score: 400},
		{ID: 3, Score: 600},
	}

	customers := []csbalancing.Entity{
		{ID: 1, Score: 300},
		{ID: 2, Score: 300},
		{ID: 3, Score: 400},
	}

	customerSuccessAway := []int{2}

	selectedCustomerSuccessIndex := csbalancing.CustomerSuccessBalancing(customerSuccess, customers, customerSuccessAway)

	fmt.Println("The ID of the CS that serves more customers is:", selectedCustomerSuccessIndex)
}
