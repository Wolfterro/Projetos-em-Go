package libs

import (
	"fmt"
	"os"
	"strconv"
	"io/ioutil"
	"path/filepath"
)

// ParseArgs parsing arguments from the command line
func ParseArgs() (string, int) {
	args := os.Args
	
	if len(args) < 2 {
		PrintErrorMessage("[Playah] Exemplo: playah <musica.mp3> [SampleRate]")
		os.Exit(1)

		return "", 0	// lol
	} else if len(args) == 2 {
		return args[1], 0
	} else {
		sampleRate, err := strconv.Atoi(args[2])
		if err != nil {
			PrintErrorMessage(err.Error())
			os.Exit(1)
		}

		return args[1], sampleRate
	}
}

// PrintErrorMessage prints a custom error message
func PrintErrorMessage(message string) {
	fmt.Println("[Playah] Ocorreu um erro ao reproduzir o arquivo .mp3!")
	fmt.Println(message)
}

// GetMP3FileContent gets the content of .mp3 file from the path
func GetMP3FileContent(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// GetFilename gets the filename from the path
func GetFilename(path string) string {
	return filepath.Base(path)
}