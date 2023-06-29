package csbalancing

import (
	"errors"
	"fmt"
	"sort"
)

const (
	minCsScore = 0
	maxCsScore = 1000
	minCsID    = 0
	maxCsID    = 10000
	maxCsCount = 1000
)

func CheckCustomerSuccessEntities(customerSuccess []Entity) []Entity {
	for _, cs := range customerSuccess {
		if cs.Score < minCsScore || cs.Score > maxCsScore {
			err := errors.New("CSs Score must be between 0 and 1000")
			fmt.Println(err)
		}
		if cs.ID < minCsID || cs.ID > maxCsID {
			err := errors.New("CSs ID must be between 0 and 10000")
			fmt.Println(err)
		}
		if cs.Score == 0 || cs.ID == 0 {
			err := errors.New("CSs cannot be empty")
			fmt.Println(err)
		}
	}

	if len(customerSuccess) >= maxCsCount {
		err := errors.New("CSs count cannot exceed 1000")
		fmt.Println(err)
	}

	return customerSuccess
}

func CheckDuplicateCSScore(customerSuccess []Entity) []Entity {
	csScore := make(map[int]bool)
	foundDuplicate := false

	for _, cs := range customerSuccess {
		if csScore[cs.Score] {
			foundDuplicate = true
			break
		}
		csScore[cs.Score] = true
	}

	if foundDuplicate {
		err := errors.New("CSs cannot have the same score")
		fmt.Println(err)
	}

	return customerSuccess
}

func FindClosestCSToCustomer(customers []Entity, customerCount map[int]int) map[int]int {
	minDiff := 0
	closestCSID := 0

	sort.SliceStable(customers, func(i, j int) bool {
		return customers[i].Score < customers[j].Score
	})

	for _, customer := range customers {
		for csID, csScore := range customerCount {
			diff := csScore - customer.Score
			if diff >= 0 && (diff < minDiff || minDiff == 0) {
				minDiff = diff
				closestCSID = csID
			}
		}

		if closestCSID != 0 {
			customerCount[closestCSID]++
		}

		minDiff = 0
		closestCSID = 0
	}

	return customerCount
}

func FindCSWithHighestCustomers(customerCount map[int]int) int {
	maxCustomers := 0

	for _, count := range customerCount {
		if count > maxCustomers {
			maxCustomers = count
		}
	}

	return maxCustomers
}
