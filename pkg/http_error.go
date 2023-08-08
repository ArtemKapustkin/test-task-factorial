package pkg

type HTTPError struct {
	err        error
	statusCode int
}

func NewHTTPError(err error, statusCode int) *HTTPError {
	return &HTTPError{err: err, statusCode: statusCode}
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

func (e *HTTPError) GetStatusCode() int {
	return e.statusCode
}
