package makosamengi

import (
	"fmt"

	"github.com/wondenge/kosa"
)

// Prefix is a helper function that will prefix some text to the given error.
// If the error is a makosamengi.Error, then it will be prefixed to each wrapped error.
// This is useful to use when appending multiple multi errors together in order to give better scoping.
func Prefix(err error, prefix string) error {
	if err == nil {
		return nil
	}

	format := fmt.Sprintf("%s {{err}}", prefix)
	switch err := err.(type) {
	case *Error:
		// Typed nils can reach here, so initialize if we are nil
		if err == nil {
			err = new(Error)
		}

		// Wrap each of the errors
		for i, e := range err.Errors {
			err.Errors[i] = kosa.Wrapf(format, e)
		}

		return err
	default:
		return kosa.Wrapf(format, err)
	}
}
