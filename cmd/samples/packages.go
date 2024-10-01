package samples

import (
	"fmt"
	"protobuf/generated/applications"
	"protobuf/generated/job"
	"protobuf/generated/soft"

	"google.golang.org/protobuf/encoding/protojson"
)

func RunPackages() {
	apps := applications.Applications{
		JobApplication: &job.Application{
			Version: "1.0",
			Name: "LinkedIn Intern",
		},
		SoftApplication: &soft.Application{
			Version: "1.0",
			Name: "NailIt",
		},
	}

	opts := protojson.MarshalOptions{
		EmitUnpopulated: true,
		EmitDefaultValues: true,
		Indent: "  ",
	}
	json, _ := opts.Marshal(&apps)
	fmt.Println(string(json))
}