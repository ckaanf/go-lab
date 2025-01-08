package conditional

func GoIfElse() {
	var a int = 10
	var b int = 20

	if a > b {
		println("a is greater than b")
	} else {
		println("b is greater than a")
	}
}

func GoIfElseIf() {
	var a int = 10
	var b int = 20

	if a > b {
		println("a is greater than b")
	} else if a < b {
		println("b is greater than a")
	} else {
		println("a is equal to b")
	}
}

func GoSwitchCaseDefault() {
	var a int = 10

	switch a {
	case 10:
		println("a is 10")
	case 20:
		println("a is 20")
	default:
		println("a is not 10 or 20")
	}
}

func GoSwitchCaseType() {
	var a interface{} = 10

	switch a.(type) {
	case int:
		println("a is int")
	case string:
		println("a is string")
	default:
		println("a is not int or string")
	}
}

func GoSwitchCaseWithOutExpression() {
	var a int = 10

	switch {
	case a == 10:
		println("a is 10")
	case a == 20:
		println("a is 20")
	default:
		println("a is not 10 or 20")
	}
}

func GoSwitchCaseFallThrough() {
	var a int = 10

	switch a {
	case 10:
		println("a is 10")
		fallthrough
	case 20:
		println("a is 20")
	default:
		println("a is not 10 or 20")
	}
}
