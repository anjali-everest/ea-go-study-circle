package quickcash

type PaytmWalletAccount struct {
	balance float64
}

func (pa *PaytmWalletAccount) WithDraw(amount float64) {
	pa.balance -= amount
}

func (pa *PaytmWalletAccount) CanWithDraw(amount float64) bool {
	if pa.balance < amount {
		return false
	}
	return true
}

func (pa *PaytmWalletAccount) GetIdentifier() string {
	return "PAYTM_WALLET_ACCOUNT"
}
