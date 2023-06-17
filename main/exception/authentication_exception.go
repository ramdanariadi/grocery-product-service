package exception

type AuthenticationException struct {
	Message string
}

func (a AuthenticationException) Error() string {
	return a.Message
}
