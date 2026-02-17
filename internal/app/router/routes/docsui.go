package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jolotech/jolo-mars/internal/docsui"
)

func RegisterDocsUI(r *gin.Engine) {
	h := docsui.NewHandler(docsui.DefaultSpec())

	r.GET("/", h.ServeHome)
	r.GET("/docs/spec.json", h.ServeSpec)
	r.GET("/docs/assets/*file", func(c *gin.Context) {
		// c.Param("file") includes leading slash
		h.ServeAsset(c)
	})
}
