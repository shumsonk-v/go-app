package helpers

import "os"

func GetEnv[K comparable, V any](name string, defaultValue V) V {
	envValue := os.Getenv(name)
	var iValue interface{} = envValue
	value, ok := iValue.(V)
	if len(envValue) == 0 || ok == false {
		return defaultValue
	}

	return value
}
