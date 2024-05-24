package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	//"net/http"
	"os"
)

var Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
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

	parseJson(fileBytes)
	//fmt.Println(storyFilePopulated)
}

func parseJson(jsonData []byte) (data map[string]Chapter) {
	err := json.Unmarshal(jsonData, &Story)
	if err != nil {
		panic(err)
	}

	// Properly parsed json file
	return Story
}

//func storyHandler()

func defaultMux(storyParsed map[string]Chapter) *http.ServeMux {
	mux := http.NewServeMux()
	for title := range storyParsed {
		mux.HandleFunc(title, func(w http.ResponseWriter, r *http.Request) {

		})
	}
	return mux
}

//func parseJson(jsonData []byte) {
//	err = json.Unmarshal(jsonData, &Story)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(Story)
//}
