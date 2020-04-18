# makosamengi

A Go (golang) package for representing a list of errors as a single error.

`makosa-mengi` is a package for Go that provides a mechanism for
representing a list of `error` values as a single `error`.

This allows a function in Go to return an `error` that might actually
be a list of errors. If the caller knows this, they can unwrap the
list and access the errors. If the caller doesn't know, the error
formats to a nice human-readable format.

`makosa-mengi` is fully compatible with the Go standard library
[errors](https://golang.org/pkg/errors/) package, including the
functions `As`, `Is`, and `Unwrap`. This provides a standardized approach
for introspecting on error values.

## Installation and Docs

Install using `go get github.com/wondenge/makosa-mengi`.

Full documentation is available at
http://godoc.org/github.com/wondenge/makosa-mengi

## Usage

`makosa-mengi` is easy to use and purposely built to be unobtrusive in
existing Go applications/libraries that may not be aware of it.

**Building a list of errors**

The `Append` function is used to create a list of errors. This function
behaves a lot like the Go built-in `append` function: it doesn't matter
if the first argument is nil, a `makosamengi.Error`, or any other `error`,
the function behaves as you would expect.

```go
var result error

if err := step1(); err != nil {
	result = makosamengi.Append(result, err)
}
if err := step2(); err != nil {
	result = makosamengi.Append(result, err)
}

return result
```

**Customizing the formatting of the errors**

By specifying a custom `ErrorFormat`, you can customize the format
of the `Error() string` function:

```go
var result *makosamengi.Error

// ... accumulate errors here, maybe using Append

if result != nil {
	result.ErrorFormat = func([]error) string {
		return "errors!"
	}
}
```

**Accessing the list of errors**

`makosamengi.Error` implements `error` so if the caller doesn't know about
makosamengi, it will work just fine. But if you're aware a makosamengi might
be returned, you can use type switches to access the list of errors:

```go
if err := something(); err != nil {
	if merr, ok := err.(*makosamengi.Error); ok {
		// Use merr.Errors
	}
}
```

You can also use the standard [`errors.Unwrap`](https://golang.org/pkg/errors/#Unwrap)
function. This will continue to unwrap into subsequent errors until none exist.

**Extracting an error**

The standard library [`errors.As`](https://golang.org/pkg/errors/#As)
function can be used directly with a makosamengi to extract a specific error:

```go
// Assume err is a makosamengi value
err := somefunc()

// We want to know if "err" has a "RichErrorType" in it and extract it.
var errRich RichErrorType
if errors.As(err, &errRich) {
	// It has it, and now errRich is populated.
}
```

**Checking for an exact error value**

Some errors are returned as exact errors such as the [`ErrNotExist`](https://golang.org/pkg/os/#pkg-variables)
error in the `os` package. You can check if this error is present by using
the standard [`errors.Is`](https://golang.org/pkg/errors/#Is) function.

```go
// Assume err is a makosamengi value
err := somefunc()
if errors.Is(err, os.ErrNotExist) {
	// err contains os.ErrNotExist
}
```

**Returning a makosamengi only if there are errors**

If you build a `makosamengi.Error`, you can use the `ErrorOrNil` function
to return an `error` implementation only if there are errors to return:

```go
var result *makosamengi.Error

// ... accumulate errors here

// Return the `error` only if errors were added to the makosamengi,
// otherwise return nil since there are no errors.
return result.ErrorOrNil()
```
