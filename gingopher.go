package gingopherjs

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	gbuild "github.com/gopherjs/gopherjs/build"
)

type (
	GinGopher struct {
		pkg string
	}
)

func New(pkg string) (*GinGopher, error) {
	return &GinGopher{
		pkg: pkg,
	}, nil
}

func (g *GinGopher) Handler(c *gin.Context) {
	f, err := ioutil.TempFile("", "gingopher")
	if err != nil {
		c.String(200, "console.error(%s);", strconv.Quote(err.Error()))
		return
	}
	f.Close()

	filename := f.Name()
	defer os.Remove(filename)

	options := &gbuild.Options{
		GOROOT:        "",
		GOPATH:        os.Getenv("GOPATH"),
		Verbose:       false,
		Quiet:         false,
		Watch:         false,
		CreateMapFile: false,
		Minify:        false,
		Color:         true,
	}

	finalPath := "/"

	// Try to deduce the package path.
	paths := strings.Split(options.GOPATH, ":")
	for _, p := range paths {
		candidatePath := path.Join(p, "src", g.pkg)
		st, e := os.Stat(candidatePath)
		if e == nil && st.IsDir() {
			finalPath = candidatePath
			break
		}
	}

	s := gbuild.NewSession(options)
	err = s.BuildDir(finalPath, g.pkg, filename)
	if err != nil {
		c.String(200, "console.error(%s);", strconv.Quote(err.Error()))
		return
	}

	http.ServeFile(c.Writer, c.Request, f.Name())
}
