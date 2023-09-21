package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	kind         Kind
	wrappedError error
	message      string
	operation    string
}

func New(op string) RichError {
	return RichError{
		operation: op,
	}
}

func (r RichError) Error() string {
	if r.message == "" {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r RichError) WithMessage(msg string) RichError {
	r.message = msg
	return r
}

func (r RichError) WithKind(k Kind) RichError {
	r.kind = k
	return r
}

func (r RichError) WithError(e error) RichError {
	r.wrappedError = e
	return r
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		r.wrappedError.Error()
	}

	return re.Message()
}

func (r RichError) Kind() Kind {
	return r.kind
}
