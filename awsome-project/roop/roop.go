package roop

func ForLoop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)
}

func ForLoopOnlyCondition() {
	sum := 0
	i := 1
	// 약간 while 느낌
	for i <= 100 {
		sum += i
		i++
	}
	println(sum)
}

func ForLoopInfinity() {
	sum := 0
	for {
		println(sum)
	}
}

// like forEach
func ForLoopRange() {
	str := "Hello, world!"
	for i, c := range str {
		println(i, c)
	}
}

func ForLoopBreakContinueGoto() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")
}

func ForLoopBreakLabel() {
	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	println("OK")
}
