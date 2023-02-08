package authorizedconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

var playerIDs = []string{}

func Parse(filename string) error {
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	var values map[string]interface{}
	err = json.Unmarshal(file, &values)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	str := fmt.Sprintf("%v", values["ids"])
    playerIDs = append(playerIDs, str)
	log.Printf("ids: %s\n", playerIDs)
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}