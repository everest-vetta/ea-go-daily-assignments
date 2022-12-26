package quickcash

type FakePrimaryAccount struct{}

func (fsa *FakePrimaryAccount) WithDraw(amount float64) {
}

func (fsa *FakePrimaryAccount) CanWithDraw(amount float64) bool {
	return true
}

func (fsa *FakePrimaryAccount) GetIdentifier() string {
	return "FAKE_PRIMARY_ACCOUNT"
}

type FakeSecondaryAccount struct{}

func (fsa *FakeSecondaryAccount) WithDraw(amount float64) {
}

func (fsa *FakeSecondaryAccount) CanWithDraw(amount float64) bool {
	return true
}

func (fsa *FakeSecondaryAccount) GetIdentifier() string {
	return "FAKE_SECONDARY_ACCOUNT"
}

type FakePrimaryAccountWithZeroBalance struct{}

func (fsa *FakePrimaryAccountWithZeroBalance) WithDraw(amount float64) {
}

func (fsa *FakePrimaryAccountWithZeroBalance) CanWithDraw(amount float64) bool {
	return false
}

func (fsa *FakePrimaryAccountWithZeroBalance) GetIdentifier() string {
	return "FAKE_ZERO_PRIMARY_ACCOUNT"
}
