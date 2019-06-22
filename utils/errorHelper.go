package utils

import "log"

// LogIfError returns the error log if present
func LogIfError(err error) {
	if err != nil {
		log.Printf("error occurred: %+v", err)
	}
}
