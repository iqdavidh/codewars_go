package kata

import (
	"strings"
)

//
//Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
//
//The output should be two capital letters with a dot separating them.
//
//It should look like this:
//
//Sam Harris => S.H
//
//Patrick Feeney => P.F

func AbbrevName(name string) string {
	//your code here

	lista := strings.Split(name, " ")
	abrev := ""

	for _, palabra := range lista {

		if abrev != "" {
			abrev += "."
		}
		abrev += strings.ToUpper(palabra[0:1])
	}

	return abrev
}