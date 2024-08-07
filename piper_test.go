package piper

import (
	"os"
	"testing"

	asset "github.com/amitybell/piper-asset"
	alan "github.com/amitybell/piper-voice-alan"
	jenny "github.com/amitybell/piper-voice-jenny"
)

func TestPiper(t *testing.T) {
	dataDir, err := os.MkdirTemp("", "ab-piper.")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dataDir)

	assets := map[string]asset.Asset{
		"jenny": jenny.Asset,
		"alan":  alan.Asset,
	}

	for name, asset := range assets {
		tts, err := New(dataDir, asset)
		if err != nil {
			t.Fatal(err)
		}

		// Test with default options
		t.Run(name+"_default", func(t *testing.T) {
			wav, err := tts.Synthesize("hello world")
			if err != nil {
				t.Fatalf("%s: %s\n", name, err)
			}
			if len(wav) < 44 {
				t.Fatalf("%s: Invalid wav file generated: len(%d)\n", name, len(wav))
			}
		})

		// Test with custom speed option
		t.Run(name+"_custom_speed", func(t *testing.T) {
			wav, err := tts.Synthesize("hello world", WithSpeed(1.2))
			if err != nil {
				t.Fatalf("%s: %s\n", name, err)
			}
			if len(wav) < 44 {
				t.Fatalf("%s: Invalid wav file generated: len(%d)\n", name, len(wav))
			}
		})

		// Test with custom noise option
		t.Run(name+"_custom_noise", func(t *testing.T) {
			wav, err := tts.Synthesize("hello world", WithNoise(0.5))
			if err != nil {
				t.Fatalf("%s: %s\n", name, err)
			}
			if len(wav) < 44 {
				t.Fatalf("%s: Invalid wav file generated: len(%d)\n", name, len(wav))
			}
		})

		// Test with both custom speed and noise options
		t.Run(name+"_custom_speed_and_noise", func(t *testing.T) {
			wav, err := tts.Synthesize("hello world", WithSpeed(1.2), WithNoise(0.5))
			if err != nil {
				t.Fatalf("%s: %s\n", name, err)
			}
			if len(wav) < 44 {
				t.Fatalf("%s: Invalid wav file generated: len(%d)\n", name, len(wav))
			}
		})
	}
}
