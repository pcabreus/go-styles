package variablenames

import "fmt"

// The general rule of thumb is that the length of a name should be proportional to the size of its scope and inversely
//proportional to the number of times that it is used within that scope. A variable created at file scope may require
//multiple words, whereas a variable scoped to a single inner block may be a single word or even just a character or
//two, to keep the code clear and avoid extraneous information.

var variableWithBigScope = "hello"

// A small scope is one in which one or two small operations are performed, say 1-7 lines.
func printer() {
	s := variableWithBigScope
	fmt.Println(s)
}

// A small scope is one in which one or two small operations are performed, say 1-7 lines.
func printerName(f, l string) {
	fmt.Printf("%s %s", f, l)
}

// A medium scope is a few small or one large operation, say 8-15 lines.
func superPrinter(firstName, lastName string) {
	fmt.Printf("Dear %s", firstName)

	// ... lot of code

	prefix := "Mr/Mrs"
	if lastName != "" {
		fmt.Printf("%s %s %s", prefix, firstName, lastName)
	}
	fmt.Printf("%s %s", prefix, firstName)

	// ... more code
}

// These numeric guidelines are not strict rules. Apply judgement based on context, clarity, and concision.
//A large scope is one or a few large operations, say 15-25 lines.
//A very large scope is anything that spans more than a page (say, more than 25 lines).
