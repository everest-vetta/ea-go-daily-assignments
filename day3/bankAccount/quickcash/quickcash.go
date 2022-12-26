package quickcash

type QuickCash struct {
	PrimaryAccount   Withdrawable
	SecondaryAccount Withdrawable
}

func (qc *QuickCash) GetCash(amount float64) (float64, string) {
	if qc.PrimaryAccount.CanWithDraw(amount) {
		qc.PrimaryAccount.WithDraw(amount)
		return amount, qc.PrimaryAccount.GetIdentifier()
	}
	qc.SecondaryAccount.WithDraw(amount)
	return amount, qc.SecondaryAccount.GetIdentifier()
}
