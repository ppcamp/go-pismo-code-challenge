package repositories

type Transactions interface {
	Create() error
	Get() error
}

type implTransactions struct{}

func NewTransactions() Transactions { return &implTransactions{} }

func (t *implTransactions) Create() error { return nil }
func (t *implTransactions) Get() error    { return nil }
