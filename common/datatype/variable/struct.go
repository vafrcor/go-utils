package variable

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
)

func CopyStruct(from interface{}, to interface{}) error {
	err := copier.Copy(to, from)
	if err != nil {
		fmt.Printf("[CopyStruct] structure copy has failed. error: %s\n", err.Error())
		return err
	}
	return nil
}

func StructToMap(v interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}

	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
