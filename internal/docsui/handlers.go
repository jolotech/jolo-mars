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

// ========== GET /docs/ ===========
func (h *Handler) ServeHome(c *gin.Context) {
	// serve the HTML page with the embedded spec
	c.Header("Content-Type", "text/html; charset=utf-8")
	// execute the template with the spec data
	_ = pageTmpl.Execute(c.Writer, h.spec)
}


// ========== GET /docs/spec.json ===========
func (h *Handler) ServeSpec(c *gin.Context) {
	// serve the OpenAPI spec as JSON
	c.Header("Content-Type", "application/json; charset=utf-8")
	// encode the spec as JSON without escaping HTML characters
	enc := json.NewEncoder(c.Writer)
	// prevent escaping of HTML characters to keep the spec readable
	enc.SetEscapeHTML(false)
	// encode the spec to the response
	_ = enc.Encode(h.spec)
}


// ========== GET /docs/assesets/:file ===========
func (h *Handler) ServeAsset(c *gin.Context) {
	// get the requested file from the URL parameter
	file := c.Param("file")
	// remove leading slash if present
	file = strings.TrimPrefix(file, "/")
	// prevent path traversal
	file = path.Clean(file)
	if strings.Contains(file, "..") {
		// if the cleaned path still contains "..", it's an invalid request
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// assets are under public/
	// construct the full path to the asset
	full := "public/" + file

	// read the asset data from the embedded filesystem
	data, err := EmbeddedAssets.ReadFile(full)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// set the appropriate content type based on the file extension
	switch {
	case strings.HasSuffix(file, ".css"):
		c.Header("Content-Type", "text/css; charset=utf-8")
	case strings.HasSuffix(file, ".js"):
		c.Header("Content-Type", "application/javascript; charset=utf-8")
	case strings.HasSuffix(file, ".svg"):
		c.Header("Content-Type", "image/svg+xml")
	case strings.HasSuffix(file, ".png"):
		c.Header("Content-Type", "image/png")
	default:
		c.Header("Content-Type", "application/octet-stream")
	}

	// cache static assets (but not HTML/spec)
	// set Cache-Control header to cache for 1 day (86400 seconds) and mark as immutable
	c.Header("Cache-Control", "public, max-age=86400, immutable")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(data)
}
