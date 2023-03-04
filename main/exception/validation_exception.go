package exception

type ValidationException struct {
	Message string
}

func (n ValidationException) Error() string {
	return n.Message
}
