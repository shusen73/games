package uiembed

import "embed"

//go:embed static/* static/**/*
var StaticFS embed.FS
