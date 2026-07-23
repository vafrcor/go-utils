package filesystem

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetImageConfig(t *testing.T) {
	t.Run("valid image", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "image.png")
		file, err := os.Create(path)
		if err != nil {
			t.Fatalf("create image: %v", err)
		}

		want := image.NewRGBA(image.Rect(0, 0, 3, 2))
		if err := png.Encode(file, want); err != nil {
			file.Close()
			t.Fatalf("encode image: %v", err)
		}
		if err := file.Close(); err != nil {
			t.Fatalf("close image: %v", err)
		}

		config, err := GetImageConfig(path)
		if err != nil {
			t.Fatalf("GetImageConfig() error = %v", err)
		}
		if config == nil {
			t.Fatal("GetImageConfig() returned nil config")
		}
		if config.Width != 3 || config.Height != 2 {
			t.Errorf("GetImageConfig() dimensions = %dx%d, want 3x2", config.Width, config.Height)
		}
		if config.ColorModel != color.NRGBAModel {
			t.Errorf("GetImageConfig() color model = %v, want %v", config.ColorModel, color.NRGBAModel)
		}
	})

	t.Run("missing file", func(t *testing.T) {
		config, err := GetImageConfig(filepath.Join(t.TempDir(), "missing.png"))
		if config != nil {
			t.Errorf("GetImageConfig() config = %#v, want nil", config)
		}
		if err == nil || !strings.Contains(err.Error(), "can't check open file") {
			t.Errorf("GetImageConfig() error = %v, want open-file error", err)
		}
	})

	t.Run("invalid image", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "invalid.png")
		if err := os.WriteFile(path, []byte("not an image"), 0o600); err != nil {
			t.Fatalf("write invalid image: %v", err)
		}

		config, err := GetImageConfig(path)
		if config != nil {
			t.Errorf("GetImageConfig() config = %#v, want nil", config)
		}
		if err == nil || !strings.Contains(err.Error(), "can't decode image") {
			t.Errorf("GetImageConfig() error = %v, want decode error", err)
		}
	})
}
