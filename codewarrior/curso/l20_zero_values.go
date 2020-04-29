package curso

/////////////////////////////////
// Types and Zero Values
// Go Playground: https://play.golang.org/p/zItROROXi64
/////////////////////////////////

import "fmt"

func main() {

	//** ZERO VALUES **//

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int                         // initialized with 0
	var price float64                     // initialized with 0.0
	var name string                       // initialized with empty string -> ""
	var done bool                         // initialized with false
	fmt.Println(value, price, name, done) // -> 0 0.0 ""  false
}
