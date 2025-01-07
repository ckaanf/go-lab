package data_type

import (
	"fmt"
)

//TIP <p>other package doesn't have func main()</p>

func BoolType() {
	var flag bool = true
	fmt.Println("bool value of flag is: ", flag)
}

func StringType() {
	var name string = "gopher"
	fmt.Println("string value of name is: ", name)
}

func IntegerType() {
	var age int = 25
	fmt.Println("int value of age is: ", age)

	var age8 int8 = 25
	fmt.Println("int8 value of age is: ", age8)

	var age16 int16 = 25
	fmt.Println("int16 value of age is: ", age16)

	var age32 int32 = 25
	fmt.Println("int32 value of age is: ", age32)

	var age64 int64 = 25
	fmt.Println("int64 value of age is: ", age64)

	var uage uint = 25
	fmt.Println("uint value of age is: ", uage)

	var uage8 uint8 = 25
	fmt.Println("uint8 value of age is: ", uage8)

	var uage16 uint16 = 25
	fmt.Println("uint16 value of age is: ", uage16)

	var uage32 uint32 = 25
	fmt.Println("uint32 value of age is: ", uage32)

	var uage64 uint64 = 25
	fmt.Println("uint64 value of age is: ", uage64)

	var uintptrAge uintptr = 25
	fmt.Println("uintptr value of age is: ", uintptrAge)
}

func FloatType() {
	var height float32 = 5.8
	fmt.Println("float value of height is: ", height)
}

func OtherType() {
	var byteType []byte = []byte("gopher")
	fmt.Println("byte value of byteType is: ", byteType)

	var runeType []rune = []rune("gopher")
	fmt.Println("rune value of runeType is: ", runeType)
}
