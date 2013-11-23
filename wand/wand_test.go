package wand

import (
	G "github.com/tHinqa/outside-graphicsmagick"
	"testing"
)

func TestInit(t *testing.T) {
	var w *MagickWand
	if w = NewMagickWand(); w == nil {
		t.Error("NewMagickWand() failed")
	}
	if w != nil {
		var s uint32
		t.Logf("%x, %s", s, GetVersion(&s))
		w.Destroy()
	}
	t.Log(GetPackageName())
	var d uint32
	t.Log(GetQuantumDepth(&d), d)
	t.Log(GetReleaseDate())
	t.Log(GetCopyright())
	t.Log(GetHomeURL())
}

func TestQuery(t *testing.T) {
	G.InitializeMagick("")
	NewMagickWand()
	var n uint32
	t.Log(n, "fonts", QueryFonts("*", &n))
	t.Log(n, "formats", QueryFormats("*", &n)) // 0 !?
}
