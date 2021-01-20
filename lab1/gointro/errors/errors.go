package errors

import "fmt"

/*
Task: Errors needed for multiwriter

You may find this blog post useful:
http://blog.golang.org/error-handling-and-go

Similar to a the Stringer interface, the error interface also defines a
method that returns a string.

type error interface {
    Error() string
}

Thus also the error type can describe itself as a string. The fmt package (and
many others) use this Error() method to print errors.

Implement the Error() method for the Errors type defined above.

The following conditions should be covered:

1. When there are no errors in the slice, it should return:

"(0 errors)"

2. When there is one error in the slice, it should return:

The error string return by the corresponding Error() method.

3. When there are two errors in the slice, it should return:

The first error + " (and 1 other error)"

4. When there are X>1 errors in the slice, it should return:

The first error + " (and X other errors)"
*/
func (m Errors) Error() string {

	j := 0
	a := make(Errors, len(m))
	if m == nil {
		return "(0 errors)"
	}
	for i, err := range m {
		if m[i] != nil {
			a[j] = err
			j++
		}
	}

	if j == 1 {
		return a[0].Error()
	} else if j == 2 {
		return a[0].Error() + " (and 1 other error)"
	} else if j > 2 {
		return a[0].Error() + fmt.Sprintf(" (and %v other errors)", j-1)
	}

	return "(0 errors)"
}
