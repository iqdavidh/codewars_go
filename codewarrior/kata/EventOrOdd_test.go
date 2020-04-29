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

	listaPrueba := [4]dataPrueba{
		{num: 1, respuestaEsperada: "Odd"},
		{num: 2, respuestaEsperada: "Even"},
		{num: -1, respuestaEsperada: "Odd"},
		{num: -2, respuestaEsperada: "Even"},
	}

	for _, itemPrueba := range listaPrueba {
		respuesta := EventOrOdd(itemPrueba.num)
		a.Equal(itemPrueba.respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, itemPrueba.respuestaEsperada))
	}
}

type dataPrueba struct {
	num               int
	respuestaEsperada string
}
