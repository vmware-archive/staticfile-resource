package api

import (
	"io/ioutil"
	"os"
)

type In struct {
}

func NewIn() In {
	return In{}
}

func (i In) WriteDataToDestination(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
