package brazil

import (
	"bytes"
	"embed"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/theapemachine/errnie/v2"
)

type File struct {
	Data *bytes.Buffer
}

func NewFile(path string) *File {
	replaced := strings.Replace(path, "~/", HomePath(), -1)

	buf, err := ioutil.ReadFile(replaced)
	errnie.Ambient().Handle(errnie.ERROR, errnie.RET, err)

	return &File{Data: bytes.NewBuffer(buf)}
}

/*
WriteIfNotExists is a specialized method to deal with embedded filesystems meant to
supply any missing dependencies no matter what.
*/
func WriteIfNotExists(path string, embedded embed.FS) {
	cfgFile := GetFileFromPrefix(path)
	slug := BuildPath(HomePath(), path)

	if _, err := os.Stat(slug); os.IsNotExist(err) {
		fs, err := embedded.Open("cfg/" + cfgFile)
		errnie.Ambient().Log(errnie.ERROR, err)

		defer fs.Close()

		buf, err := io.ReadAll(fs)
		errnie.Ambient().Log(errnie.ERROR, err)

		errnie.Ambient().Log(
			errnie.INFO,
			".amsh.yml not found in home path, writing embedded default to: "+slug,
		)

		err = ioutil.WriteFile(slug, buf, 0644)
		errnie.Ambient().Log(errnie.ERROR, err)
	}
}
