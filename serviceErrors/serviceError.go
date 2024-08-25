package serviceErrors

type ServiceError struct {
	WrappedError error
	Code         string
	Description  string
}

func (e ServiceError) Error() string {
	return e.WrappedError.Error()
}
