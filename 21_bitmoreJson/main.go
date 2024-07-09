package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int64 
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"reactjs bootcamp", 299, "learncode", "abc1234", []string{"webdev", "js"}},
		{"mern bootcamp", 299, "learncode", "def1234", []string{"webdev", "mern"}},
		{"golang bootcamp", 299, "learncode", "ghi1234", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}




func DecodeJson()  {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "React js",
			"price": 299,
			"website": "google.com",
			"tags": ["web-dev", "js"]
		}
	`)

	var lcoCourse course

	checkvalid:=json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}
	//some cases where you just want to add data to key value
	var myonlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myonlineData)
	fmt.Printf("%#v\n",myonlineData)

	for k,v := range myonlineData{
		fmt.Printf("key is %v and value is %v and type is : %T\n",k,v,v)
	}
}

