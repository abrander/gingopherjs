# gingopherjs
A gin-gonic route for developing client side javascript with gopherjs.

Can be used like this:

```go
import (
	"github.com/gin-gonic/gin"
	"github.com/abrander/gingopherjs"
)

func main() {
	router := gin.Default()
	g, _ := gingopherjs.New("github.com/me/myrepo")
	router.GET("/client.js", g.Handler)
	router.Run(":8080")
}
```

Please note that this will compile your Javascript at *every* request. It's only useful for development and should never be exposed on a public webserver.
