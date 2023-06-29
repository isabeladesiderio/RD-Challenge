package csbalancing

import (
	"errors"
	"fmt"
)

// Entity ...
type Entity struct {
	ID    int
	Score int
}

func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {

	customerSuccess = CheckCustomerSuccessEntities(customerSuccess)

	customers = CheckCustomerEntities(customers)

	customerSuccess = CheckDuplicateCSScore(customerSuccess)

	customerSuccess = RemoveCustomerSuccessAway(customerSuccess, customerSuccessAway)

	customerCount := CountCustomersPerCustomerSuccess(customerSuccess, customers)

	customerCount = FindClosestCSToCustomer(customers, customerCount)

	maxCustomers := FindCSWithHighestCustomers(customerCount)

	multipleMax := CheckMultipleMaxCustomers(customerCount, maxCustomers)
	if multipleMax {
		err := errors.New("More than one Customer Success with the same maximum number of customers")
		fmt.Println(err)
		return 0
	}

	maxCustomerSuccessID := FindMaxCustomerSuccessID(customerCount, maxCustomers)

	customerSuccessIndex := FindCustomerSuccessIndexByID(customerSuccess, maxCustomerSuccessID)
	if customerSuccessIndex == 0 {
		err := errors.New("Customer Success not found")
		fmt.Println(err)
		return 0
	}

	return customerSuccessIndex
}
