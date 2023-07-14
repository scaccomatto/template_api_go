package apperrors

type HttpError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message,omitempty"`
	Details    string `json:"details,omitempty"`
	MetaData   any    `json:"metaData,omitempty"`
}

func (e HttpError) Error() string {
	return e.Message
}
