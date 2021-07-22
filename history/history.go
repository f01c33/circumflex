package history

import (
	"clx/file"
	"os"
	"path"

	"github.com/emirpasic/gods/sets/hashset"
)

const (
	disableHistory = 0
)

type History struct {
	visitedStories *hashset.Set
	mode           int
}

func (his *History) Contains(id int) bool {
	if his.mode == disableHistory {
		return false
	}

	return his.visitedStories.Contains(id)
}

func Initialize(historyMode int) *History {
	h := &History{
		visitedStories: hashset.New(),
		mode:           historyMode,
	}

	if h.mode == disableHistory {
		return h
	}

	fullPath, dirPath, fileName := getCacheFilePaths()

	if !exists(fullPath) {
		writeToDisk(h, dirPath, fileName)

		return h
	}

	historyFileContent, readErr := os.ReadFile(fullPath)
	if readErr != nil {
		panic(readErr)
	}

	deserializationErr := h.visitedStories.FromJSON(historyFileContent)
	if deserializationErr != nil {
		panic(deserializationErr)
	}

	return h
}

func writeToDisk(h *History, dirPath string, fileName string) {
	emptyJSON, _ := h.visitedStories.ToJSON()

	err := file.WriteToFileNew(dirPath, fileName, string(emptyJSON))
	if err != nil {
		panic(err)
	}
}

func getCacheFilePaths() (string, string, string) {
	homeDir, _ := os.UserHomeDir()
	configDir := ".cache"
	circumflexDir := "circumflex"
	fileName := "history.json"

	fullPath := path.Join(homeDir, configDir, circumflexDir, fileName)
	dirPath := path.Join(homeDir, configDir, circumflexDir)

	return fullPath, dirPath, fileName
}

func exists(pathToFile string) bool {
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		return false
	}

	return true
}
