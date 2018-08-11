package Utilities

import (
	"net/url"
)

// Respond to errors
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Checks if string is a valid URL
func IsValidUrl(toCheck string) bool {
	_, err := url.ParseRequestURI(toCheck)
	if err != nil {
		return false
	} else {
		return true
	}
}
