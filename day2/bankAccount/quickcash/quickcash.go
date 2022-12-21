package quickcash

type SavingsAccount struct {
	balance float64
}

func (sa *SavingsAccount) WithDraw(amount float64) {
    sa.balance -= amount
}

func (sa *SavingsAccount) CanWithDraw(amount float64) bool {
    if(sa.balance < amount) {
        return false
    }
	return true
}

func (sa *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}

type CreditCardAccount struct {
	balance float64
}

func (ca *CreditCardAccount) WithDraw(amount float64) {
    ca.balance -= amount
}

func (ca *CreditCardAccount) CanWithDraw(amount float64) bool {
    if(ca.balance < amount) {
        return false
    }
	return true
}

func (ca *CreditCardAccount) GetIdentifier() string {
	return "CREDIT_CARD_ACCOUNT"
}

type QuickCash struct {
	SavingsAccount   Withdrawable
	CreditCardAccount Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	if qc.SavingsAccount.CanWithDraw(amount) {
		qc.SavingsAccount.WithDraw(amount)
		return amount, qc.SavingsAccount.GetIdentifier()
	}
	qc.CreditCardAccount.WithDraw(amount)
	return amount, qc.CreditCardAccount.GetIdentifier()
}
