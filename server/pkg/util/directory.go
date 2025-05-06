package util

import (
	"fmt"
	"os"
)

func CreateTmpDirectory() {
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		err = os.Mkdir("tmp", 0750)
		if err != nil {
			panic(fmt.Errorf("error creating tmp: %w", err))
		}
	}
}
