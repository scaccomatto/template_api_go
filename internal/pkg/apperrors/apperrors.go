package apperrors

type AppError struct {
	Status     string
	Message    string      `json:"message"`
	Details    string      `json:"details"`
	InnerError interface{} `json:"error"`
}

const (
	BadRequest string = "bad_request"
	NotFound   string = "not_found"
)

func (e AppError) Error() string {
	return e.Message
}

type ErrorBuilder struct {
	Err AppError
}

func Builder() *ErrorBuilder {
	return &ErrorBuilder{
		Err: AppError{},
	}
}

func (eb *ErrorBuilder) Message(msg string) *ErrorBuilder {
	eb.Err.Message = msg
	return eb
}

func (eb *ErrorBuilder) Details(dts string) *ErrorBuilder {
	eb.Err.Details = dts
	return eb
}

func (eb *ErrorBuilder) Error(err error) *ErrorBuilder {
	eb.Err.InnerError = err
	return eb
}

func (eb *ErrorBuilder) Status(status string) *ErrorBuilder {
	eb.Err.Status = status
	return eb
}

func (eb *ErrorBuilder) Build() error {
	return eb.Err
}
