package samples

import (
	"fmt"
	"protobuf/generated/comunication"
	"protobuf/generated/user"

	"google.golang.org/protobuf/encoding/protojson"
)

func RunOneOf() {
	p := user.User{
		ElectronicComunicationChannel: &user.User_InstantMessaging{
			InstantMessaging: &comunication.InstantMessaging{
				Name: "Instant Messaging",
				PhoneNumber: "111-2222-3333",
			},
		},
	}

	opts := protojson.MarshalOptions{
		Indent: "  ",
	}

	jsonBytes, _ := opts.Marshal(&p)
	fmt.Println(string(jsonBytes))
}