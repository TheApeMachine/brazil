package brazil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/theapemachine/errnie/v2"
)

/*
HomePath does its best to give the caller back the actual home path of the
current user, no matter which OS or environment they are on.
*/
func HomePath() string {
	home, err := os.UserHomeDir()
	errnie.Handles(err).With(errnie.RECV)
	return BuildPath(home)
}

func Workdir() string {
	wd, err := os.Getwd()
	errnie.Handles(err).With(errnie.RECV)
	return wd
}

func BuildPath(frags ...string) string {
	return filepath.FromSlash(strings.Join(frags, "/"))
}

func GetFileFromPrefix(prefix string) string {
	frags := strings.Split(prefix, "/")
	return frags[len(frags)-1]
}
