package tools

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"sync"
)

var onceMix sync.Once

var Mix map[string]string
func init()  {

	onceMix.Do(func() {

		//AbsolutePath := config.AbsolutePath()

		mixFile := "public/mix-manifest.json"

		data, err := ioutil.ReadFile(mixFile)

		if err != nil{
			fmt.Println("MIX FILE NOT FOUND")
		}
		err2 := json.Unmarshal(data, &Mix)
		if err2 != nil{
			fmt.Println("File parse error")
		}
	})

}
