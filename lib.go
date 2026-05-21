package cfg_parser

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

type Data struct {
	Title string
	Pair  map[string]string
}

type Config struct {
	Datas     map[string]Data
	lastTitle string
}

func NewConfig() *Config {
	return &Config{
		Datas:     make(map[string]Data),
		lastTitle: "",
	}
}

func (c *Config) Parse(content string) {
	cleanContent := strings.TrimSpace(content)
	scanner := bufio.NewScanner(strings.NewReader(cleanContent))
	line := 0

	for scanner.Scan() {
		line++
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if len(input) <= 2 {
			log.Fatalf("bad config, input: %s [LINE NO = %d]\n", input, line)
		}

		if line == 1 && input[0] != '[' {
			log.Fatalf("bad config, title missing [LINE NO = %d]\n", line)
		}

		switch input[0] {
		case '[':
			var title strings.Builder
			title.WriteString(string(input[1]))

			for i := 2; i < len(input); i++ {
				c := input[i]

				if i == len(input)-1 && c == ']' {
					break
				} else if i == len(input)-1 && c != ']' {
					log.Fatalf("bad config, title missing `]` at the end [LINE NO = %d]\n", line)
				} else {
					title.WriteString(string(c))
				}
			}

			cleanTitle := strings.TrimSpace(title.String())
			c.lastTitle = cleanTitle
		default:
			if c.lastTitle == "" {
				log.Fatalf("i don't know error, title empty ???, input: %s [LINE NO = %d]\n", input, line)
			}

			exist := strings.Contains(input, "=")
			if !exist {
				log.Fatalf("bad config, pair: %s [LINE NO = %d]\n", input, line)
			}
			pair := strings.SplitN(input, "=", 2)
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])

			if key == "" {
				log.Fatalf("bad config, pair key is empty [LINE = %d]\n", line)
			}
			if value == "" {
				log.Fatalf("bad config, pair value is empty [LINE = %d]\n", line)
			}

			if IsNum(key[0]) {
				log.Fatalf("key can't start with a number [KEY = %s | LINE = %d]\n", key, line)
			}

			if IsNum(value[0]) {
				// is num
				for i := range value {
					curr := value[i]
					if curr == '.' {
						if i >= len(value)-1 {
							log.Fatalf("number can't end with dot in a decimal number type [VALUE = %s | LINE = %d]\n", value, line)
						}
					}
				}
			}

			c.insert_pair(c.lastTitle, key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading string:", err)
	}
}

func IsNum(char byte) bool {
	return char >= '0' && char <= '9'
}

func (c *Config) insert_pair(title, key, value string) {
	data, exist := c.Datas[title]
	if !exist {
		data = Data{
			Title: title,
			Pair:  make(map[string]string),
		}
	}

	data.Pair[key] = value
	c.Datas[title] = data
}

func (c *Config) isTitleFound(title string) bool {
	_, exist := c.Datas[title]
	return exist
}

// (key, value, error)
func (c *Config) GrabPair(title string, pair_key string) (string, string, error) {
	data, err := c.GrabPairs(title)
	if err != nil {
		return "", "", err
	}

	pair_value, exist := data[pair_key]
	if !exist {
		return "", "", fmt.Errorf("pair with key %s not found", pair_key)
	}

	return pair_key, pair_value, nil
}

func (c *Config) GrabPairs(title string) (map[string]string, error) {
	data, exist := c.Datas[title]
	if !exist {
		return nil, fmt.Errorf("title %s not found", title)
	}
	return data.Pair, nil
}

// for debugging
func (c *Config) Print() {
	for _, data := range c.Datas {
		for key, value := range data.Pair {
			fmt.Printf("TITLE: %s | KEY: %s | VALUE: %s\n", data.Title, key, value)
		}
	}
}
