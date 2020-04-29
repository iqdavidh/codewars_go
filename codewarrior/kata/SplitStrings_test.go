package kata

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {

	a := assert.New(t)

	respuestaEsperada := []string{"ab", "c_"}
	respuesta := SplitStrings("abc")
	a.Equal(respuestaEsperada, respuesta, fmt.Sprintf("error , tenemos  %v esperamos %v", respuesta, respuesta))
}
