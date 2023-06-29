package csbalancing

import (
	"errors"
	"fmt"
)

const (
	minCtScore = 0
	maxCtScore = 100000
	minCtID    = 0
	maxCtID    = 1000000
	maxCtCount = 1000000
)

func CheckCustomerEntities(customers []Entity) []Entity {

	for _, ct := range customers {
		if ct.Score < minCtScore || ct.Score > maxCtScore {
			err := errors.New("Customer Score must be between 0 and 100000")
			fmt.Println(err)
		}
		if ct.ID < minCtID || ct.ID > maxCtID {
			err := errors.New("Customer ID must be between 0 and 1000000")
			fmt.Println(err)
		}
		if ct.Score == 0 || ct.ID == 0 {
			err := errors.New("Customer cannot be empty")
			fmt.Println(err)
		}
	}

	if len(customers) >= maxCtCount {
		err := errors.New("Customer count cannot exceed 1000000")
		fmt.Println(err)
	}

	return customers
}

func CountCustomersPerCustomerSuccess(customerSuccess []Entity, customers []Entity) map[int]int {
	customerCount := make(map[int]int)

	for _, customer := range customers {
		for _, customerSuccessEntity := range customerSuccess {
			if customer.Score <= customerSuccessEntity.Score {
				customerCount[customerSuccessEntity.ID]++
				break
			}
		}
	}

	return customerCount
}

func FindMaxCustomerSuccessID(customerCount map[int]int, maxCustomers int) int {
	var maxCustomerSuccessID int

	for customerSuccessID, count := range customerCount {
		if count == maxCustomers {
			maxCustomerSuccessID = customerSuccessID
			break
		}
	}

	return maxCustomerSuccessID
}

func FindCustomerSuccessIndexByID(customerSuccess []Entity, maxCustomerSuccessID int) int {
	for i, customerSuccessEntity := range customerSuccess {
		if customerSuccessEntity.ID == maxCustomerSuccessID {
			return i + 1
		}
	}
	return 0
}

func CheckMultipleMaxCustomers(customerCount map[int]int, maxCustomers int) bool {
	multipleMax := false

	for _, count := range customerCount {
		if count == maxCustomers {
			if multipleMax {
				return true
			}
			multipleMax = true
		}
	}

	return false
}
