package json

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetJSON(JSONurl string) []byte {
	jsonfile, err := http.Get(JSONurl)
	//defer jsonfile.Body.Close()
	if err != nil {
		fmt.Println(err)
		return []byte("err")
	} else {
		rtfile, _ := ioutil.ReadAll(jsonfile.Body)
		err := jsonfile.Body.Close()
		if err != nil {
			return nil
		}
		return rtfile
	}
}
