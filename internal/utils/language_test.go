package utils

import (
	"testing"
)

func TestDetectLanguageCode(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"English", "Hello world", "en"},
		{"Chinese", "你好世界", "zh"},
		{"Japanese (Mixed)", "こんにちは世界", "ja"},
		{"Korean", "안녕하세요", "ko"},
		{"Russian", "Привет мир", "ru"},
		{"Mixed Han/Latin (Defaults to ZH)", "Hello 世界", "zh"},
		// The ambiguous case mentioned by user:
		{"Japanese Pure Kanji (Misclassified as ZH)", "東京都", "zh"},
		{"Japanese with Iteration Mark (Should be JA)", "人々", "ja"},
		{"Japanese with Comma (Should be JA)", "東京、日本", "ja"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectLanguageCode(tt.text); got != tt.want {
				t.Errorf("DetectLanguageCode(%q) = %v, want %v", tt.text, got, tt.want)
			}
		})
	}
}
