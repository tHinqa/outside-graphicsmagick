package wand

import (
	G "github.com/tHinqa/outside-graphicsmagick"
	"os"
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

var seed = []byte{71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 240, 0,
	0, 255, 255, 255, 0, 0, 0, 33, 249, 4, 0, 0, 0, 0, 0,
	44, 0, 0, 0, 0, 1, 0, 1, 0, 0, 2, 2, 68, 1, 0, 59}

func TestMagickWand(t *testing.T) {
	w := NewMagickWand()
	defer w.Destroy()
	if w.ReadBlob(&seed[0], uint32(len(seed))) {
		w.Resize(200, 200, 0, 0)
		w.AddNoise(G.GaussianNoise)
		w.Edge(0)
		w.Enhance()
		w.Despeckle()
		w.Emboss(10, 2)
		w.Equalize()
		for _, f := range []string{"png" /*, "gif", "jpg", "svg", "xbm"*/} {
			t.Log(f, w.Write("_test."+f))
			if e := os.Remove("_test." + f); e != nil {
				t.Error(e)
			}
		}
	}
}
