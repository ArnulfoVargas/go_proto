package samples

import (
	"fmt"
	"protobuf/generated/basic"
)

func RunHello() {
	prot := basic.Hello{
		Name: "Arny",
	}

	fmt.Println(prot.String())
}

func RunScalar() {
	prot := basic.ScalarTypes{
		MyUint:   10,
		MyBool:   true,
		MyBytes:  []byte{'h', 'e', 'l', 'l', 'o'},
		MyString: "World",
		MySint:   10,
		MyDouble: 3.14159,
	}

	fmt.Println(prot.String())
}
