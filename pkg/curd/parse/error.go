package parse

import "fmt"

// newMethodSyntaxError creates syntaxError
func newMethodSyntaxError(methodName, errReason string) error {
	return methodSyntaxError{
		methodName: methodName,
		errReason:  errReason,
	}
}

type methodSyntaxError struct {
	methodName string
	errReason  string
}

func (err methodSyntaxError) Error() string {
	return fmt.Sprintf("method %s has syntax errors, specific reasons: %s", err.methodName, err.errReason)
}
