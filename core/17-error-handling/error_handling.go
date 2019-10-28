// Package erratum includes a solution for the "Error Handling" problem in the Go track on https://exercism.io.
package erratum

// Use opens a resource, calls Frob(input) on the result resource and then closes that resource (in all cases).
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}
	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			if frobError, ok := r.(FrobError); ok {
				resource.Defrob(frobError.defrobTag)
				err = frobError.inner
			} else {
				err = r.(error)
			}
		}
	}()
	resource.Frob(input)
	return err
}
