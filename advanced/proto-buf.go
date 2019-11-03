package advanced

import (
	"fmt"
	"log"
)
import "github.com/golang/protobuf/proto"

func ProtoBufExample() {

	//must run
	// protoc --go_out=. *.proto
	//command to compile proto file
	elliot := Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(&elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// let's go the other way and unmarshal
	// our protocol buffer into an object we can modify
	// and use
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for good measure
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())
}
