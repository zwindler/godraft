package services

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	DefaultTemplateStruct = TemplateStruct{}
	Version               string
)

func Register() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/step2", stepTwoHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(ErrTemplateParseFile)
	}

	err = t.Execute(w, DefaultTemplateStruct)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}

func stepTwoHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	var templateStruct TemplateStruct
	// TODO deal with errors
	deckFormat := r.FormValue("deckformat")
	deckStyle := r.FormValue("deckstyle")
	nnonlands := r.FormValue("nnonlands")
	nlands := r.FormValue("nlands")

	templateStruct.SetDefaults(deckFormat, deckStyle)

	if nnonlands != "" {
		err = templateStruct.CheckNNonLands(nnonlands)
		if err != nil {
			log.Warn("fail to read nnonlands")
		}
	}
	if nlands != "" {
		err = templateStruct.CheckNLands(nlands)
		if err != nil {
			log.Warn("fail to read nlands")
		}
	}

	templateStruct.White = getLandsFromForm(r.FormValue("white"))
	templateStruct.Blue = getLandsFromForm(r.FormValue("blue"))
	templateStruct.Black = getLandsFromForm(r.FormValue("black"))
	templateStruct.Red = getLandsFromForm(r.FormValue("red"))
	templateStruct.Green = getLandsFromForm(r.FormValue("green"))

	templateStruct.suggestLands()

	t, err := template.ParseFiles("templates/step2.html")
	if err != nil {
		log.Error(ErrTemplateParseFile)
	}

	err = t.Execute(w, templateStruct)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}
