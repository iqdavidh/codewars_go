package kata

//Complete the solution so that it splits the string
//into pairs of two characters. If the string contains
//an odd number of characters then it should replace the
//missing second character of the final pair with an underscore ('_').
//
//Examples:
//
//Solution("abc") //should return ["ab", "c_"]
//Solution("abcdef") //should return ["ab", "cd", "ef"]
//

func Solution(texto string) []string {

	size := len(texto)
	if size%2 > 0 {
		size++
		texto += "_"
	}

	contador := size / 2

	var respuesta = make([]string, contador)

	for i := 0; i < contador; i++ {
		respuesta[i] = texto[i*2 : i*2+2]
	}

	return respuesta
}

func main() {

}
