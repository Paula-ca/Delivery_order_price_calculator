package api

type ApiError interface {
	Error() (string, int)
}

type apiErr struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

func (e *apiErr) Error() (string, int) {
	return e.ErrorMessage, e.ErrorStatus
}

func NewApiError(message string, status int) ApiError {
	return &apiErr{message, status}
}
