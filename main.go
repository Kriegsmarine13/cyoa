package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
)

type StoryFile []struct {
	name    string
	content struct {
		title   string
		story   []string
		options string
		content struct {
			text string
			arc  string
		}
	}
}

func main() {
	fFlag := flag.String("f", "", "file name")
	flag.Parse()
	if *fFlag == "" {
		log.Fatalf("No filename provided")
	}

	file, err := os.Open(*fFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes := buf.Bytes()

	storyFilePopulated, err := parseJson(fileBytes)
}

func defaultMux(file StoryFile) *http.ServeMux {
	mux := http.NewServeMux()
	for _, story := range file {
		mux.HandleFunc(story.name, func(w http.ResponseWriter, r *http.Request) {

		})
	}
	return mux
}

func parseJson(jsonData []byte) (file StoryFile, err error) {
	err = json.Unmarshal(jsonData, &file)
	if err != nil {
		log.Fatal(err)
	}
	return file, err
}
