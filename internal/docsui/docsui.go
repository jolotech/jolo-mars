package docsui

import "embed"

//go:embed public/*
// EmbeddedAssets is the embedded filesystem containing all static assets (CSS, JS, images) 
// for the documentation UI.
var EmbeddedAssets embed.FS
