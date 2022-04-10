package exception

type NotFoundError struct {
	Message string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Message: error}
}

func (err NotFoundError) Error() string {
	return err.Message
}
