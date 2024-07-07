package utils

import (
	"log"
	"os"
	"reflect"
)

func ImplementsInterface(variable interface{}, interfaceType interface{}) bool {
	interfaceTypeValue := reflect.TypeOf(interfaceType).Elem()
	return reflect.TypeOf(variable).Implements(interfaceTypeValue)
}

func ReadEnv(key string, defaults ...string) string {
    val, exists := os.LookupEnv(key)
    if exists {
        return val
    }

    if len(defaults) > 0 {
        return defaults[0]
    }

    return ""
}

func ReadEnvOrPanic(key string) string {
    val := ReadEnv(key)

	if val == "" {
		log.Panicf("Environment variable %s is not set and no default value provided", key)
    	return "" 
	}

	return val
}