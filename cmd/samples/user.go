package samples

import (
	"fmt"
	"protobuf/generated/user"

	"google.golang.org/protobuf/encoding/protojson"
)

func RunUser() {
	prot := user.User{
		Id:       10,
		Username: "Arny",
		IsActive: false,
		Password: []byte{'1', '2', '3', '4', '5'},
		Emails: []string{
			"test@test.com",
			"no-name@test.com",
		},
		Gender: user.Gender_MALE,
	}

	fmt.Println(prot.String())
}

func RunJsonUser() {
	prot := user.User{
		Id:       10,
		Username: "Arny",
		IsActive: false,
		Password: []byte{'1', '2', '3', '4', '5'},
		Emails: []string{
			"test@test.com",
			"no-name@test.com",
		},
		Gender: user.Gender_MALE,
	}
	// Common style, base configuration
	// jsonBytes, err := protojson.Marshal(&prot)

	// Custom marshaling
	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
		EmitDefaultValues: true,
	}
	

	jsonBytes, err := opts.Marshal(&prot)
	if err != nil {
		panic("Error parsing to json")
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}

func RunProtoUser() {
	var json string = "{\"id\":10,\"username\":\"Arny\",\"is_active\":false,\"password\":\"MTIzNDU=\",\"emails\":[\"test@test.com\",\"no-name@test.com\"],\"gender\":\"MALE\"}"
	var p user.User

	protojson.Unmarshal([]byte(json), &p)

	fmt.Println(p.String())
}

func RunProtoNested() {
	var user user.User = user.User{
		Id: 0,
		Address: &user.Address{
			Street: "5th streeth",
			City: "Washington",
			Country: "USA",
		},
		Username: "Test",
		IsActive: true,
		Password: []byte{'1','2', '3', '4', '5'},
		Emails: []string{
			"test@test.com",
			"test2@test2.com",
		},
		Gender: user.Gender_MALE,
	}

	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
		EmitDefaultValues: true,
		Indent: "  ",
	}
	jsonBytes, _ := opts.Marshal(&user)
	fmt.Println(string(jsonBytes))
}

func RunProtoNestedMessages() {
	var user user.User = user.User{
		Id: 0,
		Address: &user.Address{
			Street: "5th streeth",
			City: "Washington",
			Country: "USA",
			Coordinate: &user.Address_Coords{
				Latitude: 104.2,
				Logitude: 10.204,
			},
		},
		Username: "Test",
		IsActive: true,
		Password: []byte{'1','2', '3', '4', '5'},
		Emails: []string{
			"test@test.com",
			"test2@test2.com",
		},
		Gender: user.Gender_MALE,
	}

	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
		EmitDefaultValues: true,
		Indent: "  ",
	}
	jsonBytes, _ := opts.Marshal(&user)
	fmt.Println(string(jsonBytes))
}