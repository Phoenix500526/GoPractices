package main

import(
	"log"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main(){
	test := &Student{
		Name: "Phoenix",
		Male: true,
		Scores:[]int32{99,100,98},
	}
	data, err := proto.Marshal(test)
	if err != nil{
		log.Fatal("marshaling error: ", err)
	}
	var newTest Student
	err = proto.Unmarshal(data, &newTest)
	if err != nil {
		log.Fatal("Unmarshal error: ", err)
	}
	fmt.Printf("test.name = %v, newTest.name = %v\n", test.GetName(), newTest.GetName())
	fmt.Printf("test.male = %v, newTest.male = %v\n", test.GetMale(), newTest.GetMale())
	fmt.Printf("test.scores = %v, newTest.scores = %v\n", test.GetScores(), newTest.GetScores())
}