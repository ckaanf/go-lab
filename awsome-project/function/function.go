package function

func FunctionName(parameter string) string {
	return "Hello" + parameter + "return 값은 parameter 뒤에"
}

func PointerFunction(parameter *string) string {
	return "Hello" + *parameter
}

func VariableFunction(parameter ...string) {
	for _, s := range parameter {
		println(s)
	}
}
