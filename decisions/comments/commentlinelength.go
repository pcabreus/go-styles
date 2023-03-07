package comments

// Comment line length
// Ensure that commentary is readable from source even on narrow screens.
//
// When a comment gets too long, it is recommended to wrap it into multiple single-line comments.
// When possible, aim for comments that will read well on an 80-column wide terminal, however this is not a hard cut-off;
// there is no fixed line length limit for comments in Go. The standard library, for example,
// often chooses to break a comment based on punctuation, which sometimes leaves the individual lines closer to the 60-70 character mark.
//
// There is plenty of existing code in which comments exceed 80 characters in length.
// This guidance should not be used as a justification to change such code as part of a readability review (see consistency),
// though teams are encouraged to opportunistically update comments to follow this guideline as a part of other refactors.
// The primary goal of this guideline is to ensure that all Go readability mentors make the same recommendation when and if recommendations are made.
//
// See this post from The Go Blog on documentation <https://blog.golang.org/godoc-documenting-go-code> for more on commentary.
//

// This is a comment paragraph.
// The length of individual lines doesn't matter in Godoc;
// but the choice of wrapping makes it easy to read on narrow screens.
//
// Don't worry too much about the long URL:
// https://supercalifragilisticexpialidocious.example.com:8080/Animalia/Chordata/Mammalia/Rodentia/Geomyoidea/Geomyidae/
//
// Similarly, if you have other information that is made awkward
// by too many line breaks, use your judgment and include a long line
// if it helps rather than hinders.
