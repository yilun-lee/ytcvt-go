package utils

import (
	"fmt"
	"os"
	"strings"
)

func PanicOnErr(e error) {
	if e != nil {
		panic(e)
	}
}

type ResultError struct {
	Res string
	Err error
}

type NameResultError struct {
	Name string
	ResultError
}

// -----------------------------------------------------------------------------------------------------------

func ReadFile(filePath string) []string {
	dat, err := os.ReadFile(filePath)
	PanicOnErr(err)

	var lines []string
	for _, line := range strings.Split(string(dat), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines
}

// https://stackoverflow.com/questions/40811117/equivalent-of-python-string-format-in-go
func StringReplacer(raw string, myReplaceMap map[string]string) string {
	args, i := make([]string, len(myReplaceMap)*2), 0
	for k, v := range myReplaceMap {
		args[i] = "{" + k + "}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(raw)
}
