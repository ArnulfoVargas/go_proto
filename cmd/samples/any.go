package samples

import (
	"fmt"
	"math/rand"
	"protobuf/generated/comunication"
	"protobuf/generated/user"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func RunProtoAny() {
	protouser := user.User{
		CommunicationChannel: _randomComunicationChannel(),
	}

	opts := protojson.MarshalOptions{
		EmitUnpopulated: false,
		EmitDefaultValues: false,
		Indent: "  ",
	}
	
	jsonBytes, _ := opts.Marshal(&protouser)
	fmt.Println(string(jsonBytes))
}

func RunUnmarshalAnyKnown() {
	c := comunication.SocialMedia{
		Name: "Social Media",
		UserId: 10,
	}
	var a anypb.Any

	anypb.MarshalFrom(&a, &c, proto.MarshalOptions{})

	sm := comunication.SocialMedia{}

	anypb.UnmarshalTo(&a, &sm, proto.UnmarshalOptions{})

	opts := protojson.MarshalOptions{
		Indent: "  ",
	}
	jsonBytes, _ := opts.Marshal(&sm)

	fmt.Println(string(jsonBytes))
}

func RunUnmarshalAnyUnknown() {
	c := _randomComunicationChannel()

	var p protoreflect.ProtoMessage
	p, _ = anypb.UnmarshalNew(c, proto.UnmarshalOptions{})

	opts := protojson.MarshalOptions{
		Indent: "  ",
	}

	jsonBytes, _ := opts.Marshal(p)
	fmt.Println(string(jsonBytes))
}

func RunAnyIs() {
	c := _randomComunicationChannel()
	sm := comunication.SocialMedia{}

	if c.MessageIs(&sm) {
		anypb.UnmarshalTo(c, &sm, proto.UnmarshalOptions{})
		opts := protojson.MarshalOptions{
			Indent: "  ",
		}
		jsonBytes, _ := opts.Marshal(&sm)

		fmt.Println(string(jsonBytes))
	} else {
		fmt.Println(c.MessageName().Name())
	}
}

func _randomComunicationChannel() *anypb.Any {
	randomVal := rand.Int() % 3;

	var a anypb.Any;

	switch randomVal {
	default:
		com := comunication.InstantMessaging{ Name: "instant messaging", PhoneNumber: "111-2222-3333" } 
		anypb.MarshalFrom(&a, &com, proto.MarshalOptions{})

	case 0:
		com := comunication.Paper{ Name: "Paper", Address: "elm streeth", Receiver: "test" } 
		anypb.MarshalFrom(&a, &com, proto.MarshalOptions{})
	case 1: 
		com := comunication.SocialMedia{ Name: "Social Media", UserId: 10 } 
		anypb.MarshalFrom(&a, &com, proto.MarshalOptions{})
	}

	return &a;
}