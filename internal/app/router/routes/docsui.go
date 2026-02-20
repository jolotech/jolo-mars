package routes

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/docsui"
)

func RegisterDocsUI(r *gin.Engine) {
	// serve the HTML
	h := docsui.NewHandler(docsui.DefaultSpec())
	r.GET("/", h.ServeHome)

	// serve the JSON spec
	r.GET("/docs/spec.json", h.ServeSpec)

	// serve embedded assets (css/js/svg) correctly
	sub, err := fs.Sub(docsui.EmbeddedAssets, "public")
	if err != nil {
		panic(err)
	}
	r.StaticFS("/docs/assets", http.FS(sub))
	r.GET("/favicon.ico", func(c *gin.Context) {
       b, err := docsui.EmbeddedAssets.ReadFile("public/joloFav.png")
        if err != nil {
           c.Status(404)
           return
        }
        c.Data(200, "image/png", b)
    })
}
