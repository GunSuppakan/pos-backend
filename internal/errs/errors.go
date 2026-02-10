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

func (e *AppError) WithMessage(msg string) *AppError {
	return &AppError{
		Status:     e.Status,
		HTTPStatus: e.HTTPStatus,
		Message:    msg,
		Err:        e.Err,
	}
}
