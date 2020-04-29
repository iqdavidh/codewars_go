package kata

//https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/train/go
//Create a function (or write a script in Shell)
//that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

func EventOrOdd(num int) string {

	respuesta := "Odd"
	if num%2 == 0 {
		respuesta = "Even"
	}

	return respuesta
}
