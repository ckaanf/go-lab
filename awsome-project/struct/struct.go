package _struct

import "fmt"

// 구조체 선언 필드 대문자 시작 Public 소문자 시작 private
type Student struct {
	Name  string
	Grade int
	No    int
	Score float64
}

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func StructurePrac() {
	var house House
	house.Address = "서울시 강남구 ..."
	house.Size = 28
	house.Price = 10
	house.Category = "아파트"

	fmt.Printf("%v\n", house)

	fmt.Printf(
		"주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n",
		house.Address,
		house.Size,
		house.Price,
		house.Category,
	)
}

func StructureInit() {
	// 구조체 초기화
	//var house House = House{"서울시 강남구 ...", 28, 10, "아파트"}
	//var house1 House = House{Size: 28, Category: "아파트"}
}

func EmbeddedStructure() {
	type User struct {
		Name string
		ID   string
		Age  int
	}

	type VIPUser struct {
		User     // embedded field
		VIPLevel int
		Price    int
	}

	user := User{"송하나", "hana", 23}
	vip := VIPUser{
		user,
		0,
		1000,
	}

	fmt.Printf("User: %s ID: %s Age: %d\n", vip.Name, vip.ID, vip.Age)
}
