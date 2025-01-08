package data_type

import "fmt"

func StringLiteral() {

	//TIP <p>String Back Quote(``)로 감싼 문자열은 Raw String</p>
	// java -> String rawLiteral = "this is a raw literal \\n and escape characters are not interpreted";
	rawLiteral := `this is a raw literal \n
	and escape characters are not interpreted`

	interLiteral := "this is a inter literal\n and escape characters are not interpreted"
	interLiteral2 := "this is a inter literal\n" + "and + operator is used to concatenate strings"

	fmt.Println("rawLiteral : " + rawLiteral)
	fmt.Println("interLiteral : " + interLiteral)
	fmt.Println("interLiteral2 : " + interLiteral2)
}
