package env

import (
	"fmt"
	"os"
	"strconv"
)

func panicLoad(key string) {
	panic(fmt.Sprintf("cannot load : %s", key))
}

func MustStr(key string) string {
	str := os.Getenv(key)
	if str == "" {
		panicLoad(key)
	}
	return str
}

func MustInt(key string) int {
	str := MustStr(key)
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func MustBool(key string) bool {
	str := MustStr(key)
	b, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return b
}
