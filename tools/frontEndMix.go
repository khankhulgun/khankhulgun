package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


func FrontMix(path string) map[string]string  {
	var frontEndMix map[string]string
	//AbsolutePath := config.AbsolutePath()

	mixFile := "public/"+path

	data, err := ioutil.ReadFile(mixFile)

	if err != nil{
		fmt.Println("MIX FILE NOT FOUND")
	}
	err2 := json.Unmarshal(data, &frontEndMix)
	if err2 != nil{
		fmt.Println("File parse error")
	}

	return  frontEndMix
}


func CallMix(index string, mixData map[string]string) string{
	return mixData[index]
}