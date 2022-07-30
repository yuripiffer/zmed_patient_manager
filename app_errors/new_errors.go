package app_errors

import "fmt"

func NewInvalidNameError(detail string, err error) AppError {
	return &ErrorStruct{
		Title:    "InvalidName",
		Detail:   fmt.Sprintf("%s", detail),
		Err:      err.Error(),
		ErrorKey: KeyInvalidNameError,
	}
}

func NewInternalServerError(detail string, err error) AppError {
	return &ErrorStruct{
		Title:    "Internal Server Error",
		Detail:   fmt.Sprintf("%s", detail),
		Err:      err.Error(),
		ErrorKey: KeyInternalServerError,
	}
}

func NewUnitError(detail string, err error) AppError {
	return &ErrorStruct{
		Title:    "Unit Error",
		Detail:   fmt.Sprintf("%s", detail),
		Err:      err.Error(),
		ErrorKey: KeyUnitError,
	}
}

func NewInputError(detail string, err error) AppError {
	return &ErrorStruct{
		Title:    "Input Error",
		Detail:   fmt.Sprintf("%s", detail),
		Err:      err.Error(),
		ErrorKey: keyInputError,
	}
}
