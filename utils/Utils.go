package utils

import (
	"strings"
	"os"
	"strconv"
	"fmt"
)

func Concat(a string, b string) string {
	s := []string{a, b}
	return strings.Join(s, "")
}

//TODO aprender a hacer 2 returns uno con un error y otro con el valor
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
    return i
}