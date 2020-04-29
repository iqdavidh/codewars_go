package kata

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbrev(t *testing.T) {

	a := assert.New(t)

	{
		respuestaEsperada := "S.H"
		respuesta := AbbrevName("Sam Harris")
		a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))
	}

}
