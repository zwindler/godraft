package services

import (
	"errors"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	ErrTemplateParseFile = errors.New("error while parsing template file")
	ErrTemplateExecute   = errors.New("error while execution of template file")
	EmptyCompute         = Compute{
		NCards: "40",
		NLands: "17",
		White:  "0",
		Blue:   "0",
		Black:  "0",
		Red:    "0",
		Green:  "0",
	}
)

type Compute struct {
	NCards string
	NLands string
	White  string
	Blue   string
	Black  string
	Red    string
	Green  string
}

func Register() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/compute", computeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(ErrTemplateParseFile)
	}

	err = t.Execute(w, EmptyCompute)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}

func computeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var templateStruct Compute
	// TODO check input validity

	templateStruct.NCards = r.FormValue("ncards")
	templateStruct.NLands = r.FormValue("nlands")
	templateStruct.White = r.FormValue("white")
	templateStruct.Blue = r.FormValue("blue")
	templateStruct.Black = r.FormValue("black")
	templateStruct.Red = r.FormValue("red")
	templateStruct.Green = r.FormValue("green")

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(ErrTemplateParseFile)
	}

	err = t.Execute(w, templateStruct)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}
