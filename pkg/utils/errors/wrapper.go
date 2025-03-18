package errors

import (
	"fmt"
)

type Error struct{ Base, Wrapped error }

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Base.Error(), e.Wrapped.Error())
}
