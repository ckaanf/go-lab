package data_type

func TypeConversion() {
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(u, f)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
