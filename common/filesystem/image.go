package filesystem

import (
	"fmt"
	"image"
	"os"
)

func GetImageConfig(path string) (*image.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't check open file with err %s", err)
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return nil, fmt.Errorf("can't decode image with error %s", err)
	}
	// fmt.Printf("image config: %#v\n", config)
	return &config, nil
}
