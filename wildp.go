package wildp

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// Args hold the command-line arguments, starting with the program name.
// If argument has wildcard character, wildp extends and merges it.
var Args []string

var (
	wildcard = regexp.MustCompile(`\*|\?|\[`)
	regation = regexp.MustCompile(`\[![^\]]+\]`)
)

func init() {
	if runtime.GOOS != "windows" {
		Args = os.Args
		return
	}

	Args = append(Args, os.Args[0])
	if len(os.Args) > 1 {
		Args = append(Args, expand(os.Args[1:])...)
	}
}

func expand(args []string) []string {
	var out []string
	for _, arg := range args {
		if exists(arg) || !wildcard.Match([]byte(arg)) {
			out = append(out, arg)
			continue
		}

		if regation.Match([]byte(arg)) {
			arg = strings.Replace(arg, "!", "^", 1)
		}

		matches, err := filepath.Glob(arg)
		if err != nil {
			// filepath.Glob returns ErrBadPattern only.
			out = append(out, arg)
			continue
		}

		for _, match := range matches {
			out = append(out, match)
		}
	}
	return out
}

func exists(arg string) bool {
	_, err := os.Stat(arg)
	return err == nil
}
