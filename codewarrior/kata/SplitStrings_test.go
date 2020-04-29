package kata

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	var a = assert.New(t)

	var respuestaEsperada = []string{"ab", "c_"}
	var respuesta = SplitStrings("abc")
	a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))
}
