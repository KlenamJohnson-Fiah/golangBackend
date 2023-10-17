package service

func deposit(currentBalance, depositAmount float64) float64 {
	currentBalance += depositAmount
	return currentBalance

}

func withdraw(currentBalance, withdrawalAmount float64) float64 {
	currentBalance -= withdrawalAmount
	return currentBalance

}
