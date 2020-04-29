package kata

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Expect(EvenOrOdd(1)).To(Equal("Odd"))
//Expect(EvenOrOdd(2)).To(Equal("Even"))
//Expect(EvenOrOdd(-1)).To(Equal("Odd"))
//Expect(EvenOrOdd(-2)).To(Equal("Even"))

func TestEventOrOdd(t *testing.T) {

	a := assert.New(t)

	{
		respuestaEsperada := "Odd"
		respuesta := EventOrOdd(1)
		a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))
	}

	{
		respuestaEsperada := "Even"
		respuesta := EventOrOdd(2)
		a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))
	}
}
