package operator

func ArithmeticOperator() {
	var a int = 10
	var b int = 20

	println("a + b = ", a+b)
	println("a - b = ", a-b)
	println("a * b = ", a*b)
	println("a / b = ", a/b)
	println("a % b = ", a%b)
	a++
	println("a ++ = ", a)
	b--
	println("b -- = ", b)
}

func RelationalOperator() {
	var a int = 10
	var b int = 20

	println("a == b = ", a == b)
	println("a != b = ", a != b)
	println("a > b = ", a > b)
	println("a < b = ", a < b)
	println("a >= b = ", a >= b)
	println("a <= b = ", a <= b)
}

func LogicalOperator() {
	var a bool = true
	var b bool = false

	println("a && b = ", a && b)
	println("a || b = ", a || b)
	println("!a = ", !a)
}

func BitwiseOperator() {
	var a uint = 60
	var b uint = 13

	println("a & b = ", a&b)
	println("a | b = ", a|b)
	println("a ^ b = ", a^b)
	println("a << 2 = ", a<<2)
	println("a >> 2 = ", a>>2)
}

func AssignmentOperator() {
	var a int = 10
	var b int = 20

	a = b
	println("a = b : ", a)
	a += b
	println("a += b : ", a)
	a -= b
	println("a -= b : ", a)
	a *= b
	println("a *= b : ", a)
	a /= b
	println("a /= b : ", a)
	a %= b
	println("a %= b : ", a)
	a <<= b
	println("a <<= b : ", a)
	a >>= b
	println("a >>= b : ", a)
	a &= b
	println("a &= b : ", a)
	a ^= b
	println("a ^= b : ", a)
	a |= b
	println("a |= b : ", a)
}

func PointerOperator() {
	var a int = 20
	var ip *int

	ip = &a

	println("Address of a variable: ", &a)
	println("Address stored in ip variable: ", ip)
	println("Value of *ip variable: ", *ip)
}
