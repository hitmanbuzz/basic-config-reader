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
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if len(input) <= 2 {
			log.Fatal("bad config, input:", input)
		}

		switch input[0] {
		case '[':
			var title strings.Builder
			title.WriteString(string(input[1]))

			for i := 2; i < len(input); i++ {
				c := input[i]

				if i == len(input)-1 && c == ']' {
					break
				} else if c == '\n' {
					log.Fatal("bad config, title:", input)
					break
				} else {
					title.WriteString(string(c))
				}
			}

			c.lastTitle = strings.TrimSpace(title.String())
		default:
			if c.lastTitle == "" {
				log.Fatal("i don't know error, title empty ???, input:", input)
			}

			exist := strings.Contains(input, "=")
			if !exist {
				log.Fatal("bad config, pair:", input)
			}
			pair := strings.SplitN(input, "=", 2)
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])

			c.insert_pair(c.lastTitle, key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading string:", err)
	}
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
