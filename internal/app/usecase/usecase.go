package usecase

type (
	ErrorType string
	Error     struct {
		Message string
		Cause   error
		Type    ErrorType
	}
)

var (
	ErrorTypeBadRequest    = ErrorType("bad_request")
	ErrorTypeInternalError = ErrorType("internal_error")
	ErrorTypeNotFound      = ErrorType("not_found")
)

func NewError(message string, cause error, errorType ErrorType) Error {
	return Error{
		Message: message,
		Cause:   cause,
		Type:    errorType,
	}
}

func (e Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Cause.Error()
}
