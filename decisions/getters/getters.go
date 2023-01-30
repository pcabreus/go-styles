package getters

//Function and method names should not use a Get or get prefix, unless the underlying concept uses the word “get”
//(e.g. an HTTP GET). Prefer starting the name with the noun directly, for example use Counts over GetCounts.
//
//If the function involves performing a complex computation or executing a remote call, a different word like Compute
//or Fetch can be used in place of Get, to make it clear to a reader that the function call may take time and
//could block or fail.

type user struct {
	name string
}

func (u user) Name() string {
	return u.name
}

// Bad
// func (u user) GetName() string
