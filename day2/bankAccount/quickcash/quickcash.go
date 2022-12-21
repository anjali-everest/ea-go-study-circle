package quickcash

type QuickCash struct {
	Accounts []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	for _, account := range qc.Accounts {
		if account.CanWithDraw(amount) {
			account.WithDraw(amount)
			return amount, account.GetIdentifier()
		}
	}
	return amount, "Unable to get cash"
}
