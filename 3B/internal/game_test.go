package internal

import (
	"3B/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGame(t *testing.T) {
	dataFolder := "../test/data"
	input, output := prepareTestcases(dataFolder)

	for c, in := range input {
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
	logWriter := &utils.LogWriter{}
	log.SetOutput(logWriter)
	log.SetFlags(0)

	game := NewGame(strings.Split(input, "\n"))
	game.Start()

	logs := logWriter.GetLogs()
	expected := strings.Split(output, "\n")
	actual := strings.Split(logs, "\n")
	for i, s := range expected {
		if i >= len(actual) {
			t.Fatalf("output is ended, %s missed", s)
		}
		if s != actual[i] {
			t.Fatalf("\n%s\n != \n%s\n", s, actual[i])
		}
	}
}
