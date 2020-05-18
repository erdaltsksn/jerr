package jerr

// BaseError interface exposes additional information about the error.
type BaseError interface {
	// Implements the built-in error interface.
	Error() string
}

// baseError is a simple error struct.
type baseError struct {
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

func (e *baseError) Error() string {
	msg := "{"
	if e.Message != "" {
		msg += `"message":"` + e.Message + `"`
	}
	if e.Details != "" {
		if e.Details[0] == '{' {
			msg += `,"details":` + e.Details
		} else {
			msg += `,"details":"` + e.Details + `"`
		}
	}
	msg += "}"

	return msg
}

// New returns a new `jerr.BaseError` with given values.
func New(message string) error {
	return &baseError{
		Message: escapeJSON(message),
	}
}

// Wrap adds an `err error` into `details` field of a new `jerr.BaseError`.
func Wrap(err error, message string) error {
	if err == nil {
		return New(message)
	}

	if message == "" {
		switch v := err.(type) {
		case *baseError:
			return v
		default:
			return New(v.Error())
		}
	}

	var details string
	switch d := err.(type) {
	case *baseError:
		details = d.Error()
	default:
		details = escapeJSON(d.Error())
	}

	return &baseError{
		Message: escapeJSON(message),
		Details: details,
	}
}
