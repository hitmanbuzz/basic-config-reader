package cfg_parser

import (
	"log"
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

	config := NewConfig(input)
	config.Parse()

	_, name, err := config.GrabPair("Information", "name")
	if err != nil {
		log.Fatalf("key not found: name\n")
	}
	goodName := "Music"

	if name != goodName {
		t.Errorf("should be Music, got %s", name)
	}

	_, version, err := config.GrabPair("Base", "version")
	if err != nil {
		log.Fatalf("key not found: version\n")
	}
	goodVersion := "0.5"

	if version != goodVersion {
		t.Errorf("should be 0.5, got %s", version)
	}

	// debugging
	// config.Print()
}
