package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	gtranslate "github.com/gilang-as/google-translate"

	"github.com/gorilla/mux"
	"thdtt.com/gtranslate-api/middlewares"
)

func main() {

	r := mux.NewRouter()

	fmt.Println("🚀 Serving GTranslate API")

	r.HandleFunc("/translate/{query}", Translate)

	wrpR := middlewares.NewLogger(r)
	port := "8730"

	err := http.ListenAndServe(":"+port, wrpR)
	if err != nil {
		return
	}
}

func Translate(writer http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	query := Query{}
	err := decoder.Decode(&query)
	if err != nil {
		fmt.Println("CANNOT PARSE QUERY")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
	}

	value := gtranslate.Translate{
		Text: query.Src,
		From: "vi",
		To:   "en",
	}
	translated, err := gtranslate.Translator(value)
	if err != nil {
		panic(err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
	} else {
		prettyJSON, err := json.MarshalIndent(translated, "", "\t")
		if err != nil {
			panic(err)
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)

		res := gtranslate.Translate{}
		json.Unmarshal(prettyJSON, &res)

		fmt.Println(string(prettyJSON))

		json.NewEncoder(writer).Encode(res)
	}

}

type Query struct {
	Src string `json:"src"`
}