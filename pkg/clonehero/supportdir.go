package clonehero

import (
	"log"
	"os"
)

var homedir string

func init() {
	var err error
	homedir, err = os.UserHomeDir()
	if err != nil {
		log.Fatal("unable to resolve the home directory")
	}
}
