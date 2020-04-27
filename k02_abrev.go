package katatest

import (
	"curso/codewars/codewarrior/kata"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbrev(t *testing.T) {

	var a = assert.New(t)

	{
		respuestaEsperada := "S.H"
		var respuesta = kata.AbbrevName("Sam Harris")
		a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))
	}

}
