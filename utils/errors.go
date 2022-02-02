package utils

import "fmt"

type ValidationError struct {
	err error
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation error : %v", v.err)
}

func NewValidationError(err error) error {
	return &ValidationError{err: err}
}

type InternalError struct {
	err error
}

func (v *InternalError) Error() string {
	return fmt.Sprintf("internal error : %v", v.err)
}

func NewInternalError(err error) error {
	return &InternalError{err: err}
}