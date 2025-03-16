package repositories

type Operations interface {
	Create() error
	Get() error
}

type implOperations struct{}

func NewOperations() Operations { return &implOperations{} }

func (t *implOperations) Create() error { return nil }
func (t *implOperations) Get() error    { return nil }
