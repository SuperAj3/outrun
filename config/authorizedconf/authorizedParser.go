package authorizedconf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

var AuthorizedPlayerIDs []string

func Parse(filename string) error {
	file, err := loadFile(filename)
	if err != nil {
		return err
	}
	AuthorizedPlayerIDs := []string{}
	var values map[string]interface{}
	err = json.Unmarshal(file, &values)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	str := fmt.Sprintf("%v", values["ids"])
    AuthorizedPlayerIDs = values["ids"]
	log.Printf("ids: %s\n", AuthorizedPlayerIDs)
	return nil
}

func loadFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}