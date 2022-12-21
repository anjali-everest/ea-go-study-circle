package quickcash

type CreditCardAccount struct {
	balance float64
}

func (ca *CreditCardAccount) WithDraw(amount float64) {
	ca.balance -= amount
}

func (ca *CreditCardAccount) CanWithDraw(amount float64) bool {
	if ca.balance < amount {
		return false
	}
	return true
}

func (ca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}
