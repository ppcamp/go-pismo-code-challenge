package repositories

type Account interface {
	Create() error
	Get() error
}

type implAccount struct{}

func NewAccount() Account { return &implAccount{} }

func (t *implAccount) Create() error { return nil }
func (t *implAccount) Get() error    { return nil }
