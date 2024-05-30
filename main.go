package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"html/template"
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

	storyParsed := parseJson(fileBytes)
	http.ListenAndServe(":9999", defaultMux(storyParsed))
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
		mux.HandleFunc("/"+title, func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFiles("./templates/template.html")
			if err != nil {
				log.Fatal(err)
			}

			err = tmpl.Execute(w, storyParsed[title])
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
