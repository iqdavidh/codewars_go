//https://www.codewars.com/kata/56606694ec01347ce800001b/train/go

//Implement a method that accepts 3 integer values a, b, c.
//The method should return true if a triangle can be built
//with the sides of given length and false in any other case.
//
//(In this case, all triangles must have surface greater than 0 to be accepted).

package kata

func IsTriangle(a int, b int, c int) bool {

	//validacion
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	//caso equilatero
	if a == b && b == c {
		return true
	}
	//https://es.wikipedia.org/wiki/Desigualdad_triangular
	//desigualdad del trianglo

	isValid := ((a + b) > c) && ((b + c) > a) && ((c + a) > b)

	return isValid

}
