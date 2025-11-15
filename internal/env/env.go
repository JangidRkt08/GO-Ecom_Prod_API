package env

import "os"

func GetString(key string) string {
	val :=os.Getenv(key)
	return val
}