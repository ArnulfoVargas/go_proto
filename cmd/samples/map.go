package samples

import (
	"fmt"
	"protobuf/generated/user"

	"google.golang.org/protobuf/encoding/protojson"
)

func RunMap() {
	p := user.User{
		SkillRating: map[string]uint32{
			"swimming" : 4,
			"running" : 3,
		},
	}

	opts := protojson.MarshalOptions{
		Indent: "  ",
	}

	jsonBytes, _ := opts.Marshal(&p)
	fmt.Println(string(jsonBytes))
}