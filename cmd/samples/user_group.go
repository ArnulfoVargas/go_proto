package samples

import (
	"fmt"
	"protobuf/generated/user"
	"protobuf/generated/user_group"

	"google.golang.org/protobuf/encoding/protojson"
)

func RunUserGroup() {
	userGroup := user_group.UserGroup{
		GroupId: 255,
		GroupName: "Test Group",
		Roles: []user_group.Roles{
			user_group.Roles_ADMIN,
			user_group.Roles_ADMIN,
			user_group.Roles_USER,
			user_group.Roles_USER,
			user_group.Roles_USER,
		},
		Users: []*user.User{
			{
				Id: 10,
				Username: "test",
				IsActive: true,
				Emails: []string{
					"test@test.com",
					"test2@test.com",
				},
				Address: &user.Address{
					Street: "Elm Street",
					City: "Gotham",
					Country: "USA",
				},
				Gender: user.Gender_FEMALE,
				Password: []byte{'1','2','3', '4', '5'},
			},
			{
				Id: 10,
				Username: "test2",
				IsActive: true,
				Emails: []string{
					"test@test.com",
					"test2@test.com",
				},
				Address: &user.Address{
					Street: "Groove Street",
					City: "Los Santos",
					Country: "USA",
				},
				Gender: user.Gender_MALE,
				Password: []byte{'h','e','l', 'l', 'o'},
			},
		},
		Description: "This is an amazing group",
	}

	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
		EmitDefaultValues: true,
		Indent: "  ",
	}

	jsonBytes, _ := opts.Marshal(&userGroup);

	fmt.Println(string(jsonBytes))
}