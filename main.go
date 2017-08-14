package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
)
import (
	"fmt"
)

//type User struct {
//	Name           string   `json:"name"`
//	Resourcegroups []string `json:"resourcegroups"`
//	Roles          []string `json:"roles"`
//	test           string   `json:"test"`
//}

type User struct {
	Name string `json:"name"`
	Test string `json:"test"`
}

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("dest:" + string(sDec))
	fmt.Println()

	var test = "emhhbmdyb0B3ZHpoMDFAdW5kZWZpbmVkOjEyMzQ1Ng=="
	sDec1, _ := b64.StdEncoding.DecodeString(test)
	fmt.Println("dest:" + string(sDec1))

	//	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	//	fmt.Println(uEnc)
	//	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	//	fmt.Println(string(uDec))

	//	var user = User{
	//		test: "hello", Name: "zhangro", Roles: []string{"ROLE1", "ROLE2"}}

	user := User{"hello", "zhangro"}

	data, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatalf("JSON ", err)
	}
	fmt.Printf("%s \n", data)
}
