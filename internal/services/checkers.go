package handlers

type implCheckers struct{}

// CheckHealth is a health check function that returns nil if the service is
// healthy.
func (t *implCheckers) CheckHealth() error { return nil }
