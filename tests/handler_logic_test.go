// tests/handler_logic_test.go
package tests

import (
	"strings"
	"testing"
)

func TestTitleValidation(t *testing.T) {
	tests := []struct {
		name  string
		title string
		want  bool
	}{
		{"Valid title", "Нормальный чат", true},
		{"Empty title", "", false},
		{"Only spaces", "   ", false},
		{"Too long", strings.Repeat("a", 201), false},
		{"Max length", strings.Repeat("a", 200), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			title := strings.TrimSpace(tt.title)
			isValid := len(title) > 0 && len(title) <= 200

			if isValid != tt.want {
				t.Errorf("Title '%s': got %v, want %v", tt.title, isValid, tt.want)
			}
		})
	}
}

func TestTextValidation(t *testing.T) {
	tests := []struct {
		name string
		text string
		want bool
	}{
		{"Valid text", "Сообщение", true},
		{"Empty text", "", false},
		{"Max length", strings.Repeat("a", 5000), true},
		{"Too long", strings.Repeat("a", 5001), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := len(tt.text) > 0 && len(tt.text) <= 5000

			if isValid != tt.want {
				t.Errorf("Text validation failed: got %v, want %v", isValid, tt.want)
			}
		})
	}
}

func TestLimitValidation(t *testing.T) {
	tests := []struct {
		limit int
		want  int
	}{
		{0, 20},    // дефолт
		{10, 10},   // нормальный
		{100, 100}, // максимум
		{150, 100}, // больше максимума -> максимум
		{-5, 20},   // отрицательный -> дефолт
	}

	for _, tt := range tests {
		limit := tt.limit
		if limit <= 0 {
			limit = 20
		}
		if limit > 100 {
			limit = 100
		}

		if limit != tt.want {
			t.Errorf("Limit %d: got %d, want %d", tt.limit, limit, tt.want)
		}
	}
}
