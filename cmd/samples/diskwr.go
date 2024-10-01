package samples

import (
	"fmt"
	"os"
	"path"
	"protobuf/generated/comunication"
	"protobuf/generated/user"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func RunWriteUser() {
	p := user.User{
		Id: 1,
		Username: "Test",
		IsActive: true,
		Password: []byte("hello wordl"),
		Emails: []string{
			"test@test.com",
			"test2@test2.com",
		},
		Gender: user.Gender_MALE,
		Address: &user.Address{
			Street: "Elm Street",
			City: "New York",
			Country: "USA",
			Coordinate: &user.Address_Coords{
				Latitude: 20.4513,
				Logitude: 10.3452,
			},
		},
		CommunicationChannel: _randomComunicationChannel(),
		ElectronicComunicationChannel: &user.User_SocialMedia{
			SocialMedia: &comunication.SocialMedia{
				Name: "Facebook",
				UserId: 45,
			},
		},
		SkillRating: map[string]uint32{
			"swimming" 	: 3,
			"running"	: 4,
			"coding"	: 5,
		},
	}

	data, _ := proto.Marshal(&p)
	cwd,_ := os.Getwd()
	var splitcwd []string = strings.Split(cwd, "\\")
	var basePath string

	for _, p := range(splitcwd) {
		basePath = path.Join(basePath, p)
	}

	dir := path.Join(basePath, "proto_bins", "user.bin")
	os.WriteFile(dir, data, 0644)
}

func RunReadUser() {
	var p user.User

	cwd,_ := os.Getwd()
	var splitcwd []string = strings.Split(cwd, "\\")
	var basePath string

	for _, p := range(splitcwd) {
		basePath = path.Join(basePath, p)
	}

	dir := path.Join(basePath, "proto_bins", "user.bin")

	data, _ := os.ReadFile(dir)

	proto.Unmarshal(data, &p)

	opts := protojson.MarshalOptions{
		Indent: "  ",
	}

	jsonBytes, _ := opts.Marshal(&p)
	fmt.Println(string(jsonBytes))
}