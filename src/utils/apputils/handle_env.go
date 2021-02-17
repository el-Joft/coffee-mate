package apputils

import (
	"os"
)

// GetEnv ->
func GetEnv(key string) string {
	envdata, found := os.LookupEnv(key)
	if !found {
		envdata = key
	}
	return envdata
}
