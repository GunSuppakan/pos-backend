package errs

import "net/http"

var (
	ErrBadRequest = &AppError{
		Status:     400,
		Message:    "Bad request",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrUnauthorized = &AppError{
		Status:     401,
		Message:    "Unauthorized",
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrForbidden = &AppError{
		Status:     403,
		Message:    "Forbidden",
		HTTPStatus: http.StatusForbidden,
	}

	ErrNotFound = &AppError{
		Status:     404,
		Message:    "Resource not found",
		HTTPStatus: http.StatusNotFound,
	}

	ErrNotFoundCategory = &AppError{
		Status:     404,
		Message:    "Category not found",
		HTTPStatus: http.StatusNotFound,
	}

	ErrConflict = &AppError{
		Status:     409,
		Message:    "Resource already exists",
		HTTPStatus: http.StatusConflict,
	}

	ErrInternal = &AppError{
		Status:     500,
		Message:    "Internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}
)
