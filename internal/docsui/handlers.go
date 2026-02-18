package docsui

import (
	"encoding/json"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	spec DocSpec
}

func NewHandler(spec DocSpec) *Handler {
	return &Handler{spec: spec}
}

// // GET /
// func (h *Handler) ServeHome(c *gin.Context) {
// 	c.Header("Content-Type", "text/html; charset=utf-8")
// 	_ = pageTmpl.Execute(c.Writer, h.spec)
// }

// // GET /docs/spec.json
// func (h *Handler) ServeSpec(c *gin.Context) {
// 	c.Header("Content-Type", "application/json; charset=utf-8")
// 	enc := json.NewEncoder(c.Writer)
// 	enc.SetEscapeHTML(false)
// 	_ = enc.Encode(h.spec)
// }


func (h *Handler) ServeHome(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	_ = pageTmpl.Execute(c.Writer, h.spec)
}

func (h *Handler) ServeSpec(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(c.Writer)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(h.spec)
}

// GET /docs/assets/:file
func (h *Handler) ServeAsset(c *gin.Context) {
	file := c.Param("file")
	file = strings.TrimPrefix(file, "/")
	// prevent path traversal
	file = path.Clean(file)
	if strings.Contains(file, "..") {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// assets are under public/
	full := "public/" + file

	data, err := EmbeddedAssets.ReadFile(full)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	switch {
	case strings.HasSuffix(file, ".css"):
		c.Header("Content-Type", "text/css; charset=utf-8")
	case strings.HasSuffix(file, ".js"):
		c.Header("Content-Type", "application/javascript; charset=utf-8")
	case strings.HasSuffix(file, ".svg"):
		c.Header("Content-Type", "image/svg+xml")
	default:
		c.Header("Content-Type", "application/octet-stream")
	}

	// cache static assets (but not HTML/spec)
	c.Header("Cache-Control", "public, max-age=86400, immutable")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(data)
}
