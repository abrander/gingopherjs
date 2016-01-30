package gingopherjs

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

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

func (p *GinGopher) Handler(c *gin.Context) {
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

	// FIXME: This will not work with multiple GOPATH's
	path := path.Join(os.Getenv("GOPATH"), "src", p.pkg)

	s := gbuild.NewSession(options)
	err = s.BuildDir(path, p.pkg, filename)
	if err != nil {
		c.String(200, "console.error(%s);", strconv.Quote(err.Error()))
		return
	}

	http.ServeFile(c.Writer, c.Request, f.Name())
}
