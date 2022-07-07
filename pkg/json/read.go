package json

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type LeziapiJson struct {
	Txt    string `json:"txt"`
	Author string `json:"author"`
}

var GetNew []byte
var Rd []byte

func InitJSON() int {
	GetNew = GetJSON("https://cdn.jsdelivr.net/gh/fuckur-mom/QssbAPI_Data@latest/data.json")
	err := ioutil.WriteFile("data.json", GetNew, 0644)
	if err != nil {
		return 0
	}
	Rd, _ := os.Open("data.json")
	fd := bufio.NewReader(Rd)
	count := 0
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return count - 2
}

func GetTxt(id int) (string, string) {
	var LeziapiJsonback LeziapiJson
	err := json.Unmarshal(Rd, &LeziapiJsonback)
	if err != nil {
		fmt.Println(err)
	}
	var GetAuthor string = string(LeziapiJsonback.Author[id])
	var GetTxt string = string(LeziapiJsonback.Txt[id])
	fmt.Println(GetAuthor)
	fmt.Println(GetTxt)
	return GetTxt, GetAuthor
}
