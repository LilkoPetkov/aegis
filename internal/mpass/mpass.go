package mpass

import (
	"log"
	"os"
)

func GetMasterPass() []byte {
	aegisMasterPass := os.Getenv("AEGIS_MASTER_PASS")

	if aegisMasterPass == "" {
		log.Fatalln("AEGIS_MASTER_PASS env variable is missing. Please make sure to set it up")
	}

	return []byte(aegisMasterPass)
}
