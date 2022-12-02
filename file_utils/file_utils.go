package file_utils

import (
	"fmt"
	"os"
)

func InputsFromFile(filename string) (*os.File, func(), error) {
	inputs, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", filename, err)
	}
	closeFile := func() {
		inputs.Close()
	}
	return inputs, closeFile, nil
}
