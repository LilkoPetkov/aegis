package mpass

import (
	"log"
	"os"
)

// GetMasterPass retrieves the master password from the environment variable AEGIS_MASTER_PASS.
// It terminates the application if the environment variable is not set.
//
// Returns:
//
//	The master password as a byte slice.
func GetMasterPass() []byte {
	aegisMasterPass := os.Getenv("AEGIS_MASTER_PASS")

	if aegisMasterPass == "" {
		log.Fatalln("AEGIS_MASTER_PASS env variable is missing. Please make sure to set it up")
	}

	return []byte(aegisMasterPass)
}
