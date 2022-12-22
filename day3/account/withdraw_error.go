package account

import "fmt"

type WithdrawError struct {
	currentBalance float64
	lackingBalance float64
}

func (w WithdrawError) Error() string {
	return fmt.Sprintf("Withdraw is not possible with current balance %f. Account is lacking balance of %f",
		w.currentBalance, w.lackingBalance)
}
