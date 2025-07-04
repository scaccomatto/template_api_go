package apperr

type StatusError struct {
	Status            int    `json:"status"`
	ExternalClientMsg string `json:"client_msg,omitempty"` // sometimes we do not want to expose to the client some error details
	Err               error  `json:"err"`
}

func New(err error, status int, externalMsg ...string) StatusError {
	e := StatusError{
		Status: status,
		Err:    err,
	}
	if externalMsg != nil && len(externalMsg) > 0 {
		for _, msg := range externalMsg {
			e.ExternalClientMsg += msg + " "
		}
	}
	return e
}

func (e StatusError) Error() string {
	var msg string
	if e.Err != nil {
		msg = e.Err.Error()
	}

	if e.ExternalClientMsg != "" {
		msg = e.ExternalClientMsg
	}
	return msg
}
