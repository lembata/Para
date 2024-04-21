package ui

import (
	"io/fs"
	"embed"
)


//go:embed app/dist
var uiBox embed.FS
var UIBox fs.FS

func init(){
	var err error
	UIBox, err = fs.Sub(uiBox, "app/dist")
	if err != nil {
		panic(err)
	}
}

type faviconProvider struct{}

var FaviconProvider = faviconProvider{}

func (p *faviconProvider) GetFavicon() []byte {
	return p.GetFaviconIco()
}

func (p *faviconProvider) GetFaviconIco() []byte {
	ret, _ := fs.ReadFile(UIBox, "favicon.ico")
	return ret
}
