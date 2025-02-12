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

func ArrayLoop() {
	//for i
	nums := [...]int{10, 20, 30, 40, 50} // [5]int{10,20,30...}과 같음
	nums[2] = 300
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	// for range
	var t [5]float64 = [5]float64{24.0, 25.9, 26.8, 27.9, 28.2}
	for i, v := range t {
		println(i, v)
	}

	// index 생략 value만 궁금할 때
	for _, v := range t {
		println(v)
	}
}
