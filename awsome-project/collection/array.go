package collection

func ArrayPrac() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	// 배열 초기화
	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3}

	println(a1[1])
	println(a2[1])

	// 다차원 배열
	var multiArr [1][2][3]int
	multiArr[0][1][2] = 10

	// 다차원 배열의 초기화
	var multiArr1 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	println(multiArr1[1][2])
}
