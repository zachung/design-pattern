package internal

import (
	"2-B/internal/contract"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestGame(t *testing.T) {
	dataFolder := "../test/data"
	input, output := prepareTestcases(dataFolder)

	for c, in := range input {
		if c != "straight" {
			continue
		}
		t.Run(c, func(t *testing.T) {
			runCase(t, in, output[c])
		})
	}
}

func prepareTestcases(dataFolder string) (map[string]string, map[string]string) {
	entries, err := os.ReadDir(dataFolder)
	if err != nil {
		panic(err)
	}
	input := make(map[string]string)
	output := make(map[string]string)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		fileName := entry.Name()
		file, err := os.ReadFile(filepath.Join(dataFolder, fileName))
		if err != nil {
			panic(err)
		}
		suffix := filepath.Ext(fileName)
		onlyName := strings.TrimSuffix(fileName, suffix)
		if suffix == ".in" {
			input[onlyName] = string(file)
		} else {
			output[onlyName] = string(file)
		}
	}
	return input, output
}

func runCase(t *testing.T, input, output string) {
	logWriter := new(strings.Builder)
	log.SetOutput(logWriter)
	log.SetFlags(0)

	inputLines := strings.Split(input, "\n")
	// 準備牌堆
	deckCardStrings := strings.Split(inputLines[0], " ")
	var cards []contract.Card
	for _, cardString := range deckCardStrings {
		r := regexp.MustCompile(`(.*)\[(.*)]`)
		match := r.FindStringSubmatch(cardString)
		cards = append(cards, contract.Card{contract.Suit(match[1]), match[2]})
	}
	players := new([4]*Player)
	for i, s := range inputLines[1:5] {
		players[i] = NewPlayer(s)
	}
	big2 := NewBig2()
	big2.SetDeck(NewDeck(cards))
	big2.SetPlayers(*players)
	big2.Start(inputCh(inputLines[5:]))

	logs := logWriter.String()
	expected := strings.Split(output, "\n")
	actual := strings.Split(logs, "\n")
	for i, s := range expected {
		if i >= len(actual) {
			t.Fatalf("output is ended, %s missed", s)
		}
		if s != actual[i] {
			t.Fatalf("line %d\n%s\n != \n%s\n", i+1, s, actual[i])
		}
	}
}

func inputCh(input []string) chan string {
	ch := make(chan string)
	go func() {
		for _, s := range input {
			ch <- s
		}
		close(ch)
	}()
	return ch
}
