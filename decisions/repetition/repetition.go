package repetition

import (
	"fmt"
	"strconv"
)

// A piece of Go source code should avoid unnecessary repetition. One common source of this is repetitive names,
//which often include unnecessary words or repeat their context or type. Code itself can also be unnecessarily
//repetitive if the same or a similar code segment appears multiple times in close proximity.

// * Package vs. exported symbol name
//When naming exported symbols, the name of the package is always visible outside your package, so redundant
//information between the two should be reduced or eliminated. If a package exports only one type and it is named
//after the package itself, the canonical name for the constructor is New if one is required.

//Repetitive Name 						-> Better Name
//widget.NewWidget 						-> widget.New
//widget.NewWidgetWithName 				-> widget.NewWithName
//db.LoadFromDatabase 					-> db.Load
//goatteleportutil.CountGoatsTeleported -> gtutil.CountGoatsTeleported or goatteleport.Count
//myteampb.MyTeamMethodRequest 			-> mtpb.MyTeamMethodRequest or myteampb.MethodRequest

type User struct {
	name string
	role string
}

func New() User                 { return User{} }
func NewWithName(s string) User { return User{name: s} }

// * Variable name vs. type
//The compiler always knows the type of a variable, and in most cases it is also clear to the reader what type a
//variable is by how it is used. It is only necessary to clarify the type of a variable if its value appears
//twice in the same scope.
//Repetitive Name 				-> Better Name
//var numUsers int 				-> var users int
//var nameString string			-> var name string
//var primaryProject *Project 	-> var primary *Project
var users = 123
var name = "hello worlds"
var admin = User{name: "john", role: "admin"}

// * Conversion type
//If the value appears in multiple forms, this can be clarified either with an extra word like `raw` and `parsed` or
//with the underlying representation:

func (u User) limitations() bool {
	limitRaw := `true`
	limitParsed, err := strconv.ParseBool(limitRaw)
	if err != nil {
		return false
	}
	limitStr := `si`
	limit := limitStr == "si" && limitParsed

	return limit
}

// * External context vs. local names
//Names that include information from their surrounding context often create extra noise without benefit. The
//package name, method name, type name, function name, import path, and even filename can all provide context that
//automatically qualifies all names within.

// Report
// In package "ads/targeting/revenue/reporting"
//type AdsTargetingRevenueReport struct{}
//func (u *User) UserName() string
//Bad: The full name is implicit in the package path and does not add value
type Report struct{}

func (u *User) Name() string {
	return u.name
}

// Connection
// In package "sqldb"
//type DBConnection struct{}
//Bad: DB is implicit in the package
type Connection struct{}

// Process
// In package "ads/targeting"
//func (c *Connection) Process(in *User) *Report {
//	inNameOfUser := in.Name()
//}
//Bad: is the context is just the id
func (c *Connection) Process(in *User) *User {
	name := in.Name()

	return &User{name: name}
}

func (c *Connection) Load(q string, limit *int64) error {
	return nil
}

//Repetition should generally be evaluated in the context of the user of the symbol, rather than in isolation.
//For example, the following code has lots of names that may be fine in some circumstances, but redundant in context:

// UserCount
//func (db *DBConnection) UserCount() (userCount int, err error) {
//	var userCountInt64 int64
//	if dbLoadError := db.LoadFromDatabase("count(distinct users)", &userCountInt64); dbLoadError != nil {
//		return 0, fmt.Errorf("failed to load user count: %s", dbLoadError)
//	}
//	userCount = int(userCountInt64)
//	return userCount, nil
//}
// Bad: Instead, information about names that are clear from context or usage can often be omitted:
func (c *Connection) UserCount() (int, error) {
	var count int64
	if err := c.Load("count(distinct users)", &count); err != nil {
		return 0, fmt.Errorf("failed to load user count: %s", err)
	}
	return int(count), nil
}
