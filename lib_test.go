package cfg_parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := `
		[Information]
		name = Game

		[Base]
		version = 0.3
		total = 5

		[Config]
		hp = 100
		kill = 69

		[Base]
		version = 0.5

		[Information]
		name = Music
	`

	config := NewConfig()
	config.Parse(input)

	_, name, _ := config.GrabPair("Information", "name")
	goodName := "Music"

	if name != goodName {
		t.Errorf("should be Music, got %s", name)
	}

	_, version, _ := config.GrabPair("Base", "version")
	goodVersion := "0.5"

	if version != goodVersion {
		t.Errorf("should be 0.5, got %s", version)
	}

	// debugging
	// config.Print()
}

func TestConfig_Parse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		content string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConfig()
			c.Parse(tt.content)
		})
	}
}
