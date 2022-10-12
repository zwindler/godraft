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
	nonLands := r.FormValue("nonlands")
	lands := r.FormValue("lands")

	templateStruct.SetDefaults(deckFormat, deckStyle)

	if nonLands != "" {
		err = templateStruct.CheckNonLands(nonLands)
		if err != nil {
			log.Warn("fail to read nonLands")
		}
	}
	if lands != "" {
		err = templateStruct.CheckLands(lands)
		if err != nil {
			log.Warn("fail to read lands")
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
