package csbalancing

func RemoveCustomerSuccessAway(customerSuccess []Entity, customerSuccessAway []int) []Entity {
	for _, customerSuccessID := range customerSuccessAway {
		for i, customerSuccessEntity := range customerSuccess {
			if customerSuccessEntity.ID == customerSuccessID {
				customerSuccess = append(customerSuccess[:i], customerSuccess[i+1:]...)
				break
			}
		}
	}
	return customerSuccess
}
