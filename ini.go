// Package ini provides functions for parsing INI configuration files.
package ini

import (
	"regexp"
	"strings"
)

var (
	sectionRegex = regexp.MustCompile(`^\[(.*)\]$`)
	assignRegex  = regexp.MustCompile(`^([^=]+)=(.*)$`)
)


type File map[string]Section

// A Section represents a single section of an INI file.
type Section map[string]string

// Returns a named Section. A Section will be created if one does not already exist for the given name.
// Loads and returns an INI File from a file on disk.
func LoadFile(Str string) (File) {
	StrArray := strings.Split(Str, "^")
	var fileone File
	fileone = make(map[string]Section)
	for a := 0; a < len(StrArray); a++ {
		StrArrayOne := strings.Split(StrArray[a], ":")
		StrArrayTwo := strings.Split(StrArrayOne[1], ",")
		var seone Section //定义一个section
		seone = make(map[string]string)
		for b := 0; b < len(StrArrayTwo); b++ {
			StrArrayThree := strings.Split(StrArrayTwo[b], "=")
			seone[StrArrayThree[0]] = StrArrayThree[1]
		}
		fileone[StrArrayOne[0]] = seone
	}
	return fileone
}
