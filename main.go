package main

import (
	"NomadCoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const PORT string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    *[]blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	teml := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"HOME", blockchain.GetBlockchain().AddBlock()}
	teml.Execute(rw, data)

}

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
