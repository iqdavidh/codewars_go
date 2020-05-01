package kata

//You get an array of numbers, return the sum of all of the positives ones.
//
//Example [1,-4,7,12] => 1 + 7 + 12 = 20
//
//Note: if there is nothing to sum, the sum is default to 0.

func filtrar(lista []int, fnFiltro func(int) bool) []int {

	ret := make([]int, 0, len(lista))

	for _, item := range lista {
		if fnFiltro(item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func fnFiltroValorPositivos(i int) bool {
	return i > 0
}

func SumOfPositive(numbers []int) int {

	listaFiltrada := filtrar(numbers, fnFiltroValorPositivos)
	suma := 0
	for _, item := range listaFiltrada {
		suma += item
	}
	return suma
}

/*
 Expect(PositiveSum([]int{1, 2, 3, 4, 5})).To(Equal(15))
    Expect(PositiveSum([]int{1, -2, 3, 4, 5})).To(Equal(13))
    Expect(PositiveSum([]int{})).To(Equal(0))
    Expect(PositiveSum([]int{-1, -2, -3, -4, -5})).To(Equal(0))
    Expect(PositiveSum([]int{-1, 2, 3, 4, -5})).To(Equal(9))
*/
