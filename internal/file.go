package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

func RemoveFileOrDie(filename string) {
	err := os.Remove(filename)

	if err != nil {
		logrus.Fatal(err)
	}
}
