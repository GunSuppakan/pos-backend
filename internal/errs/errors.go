package errs

type AppError struct {
	Message    string
	Status     int
	HTTPStatus int
	Err        error
}

func (e *AppError) Error() string {
	return e.Message
}
