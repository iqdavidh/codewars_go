package katatest

import (
	"curso/codewars/codewarrior/kata"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	var a = assert.New(t)

	var respuestaEsperada = []string{"ab", "c_"}
	var respuesta = kata.Solution("abc")
	a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))

	//
	//// assert equality
	//assert.Equal(123, 123, "they should be equal")
	//
	//// assert inequality
	//assert.NotEqual(123, 456, "they should not be equal")

	// assert for nil (good for errors)
	//assert.Nil(object)

	// assert for not nil (good when you expect something)
	//if assert.NotNil(object) {
	//
	//	// now we know that object isn't nil, we are safe to make
	//	// further assertions without causing any errors
	//	assert.Equal("Something", object.Value)
	//}
}
