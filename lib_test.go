package cfg_parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := `
		[Information]
		name = Game

		[Base]
		version = 0.1
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
		t.Errorf("should be 0.5, got %s", goodVersion)
	}
}
