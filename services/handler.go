package services

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	ErrTemplateParseFile  = errors.New("error while parsing template file")
	ErrTemplateExecute    = errors.New("error while execution of template file")
	ErrInvalidNCards      = errors.New("invalid NCards value")
	ErrInvalidNLands      = errors.New("invalid NLands value")
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

	templateStruct.SetDefaults(deckFormat, deckStyle)

	// TODO deal with errors
	templateStruct.NCards, err = CheckNCards(r.FormValue("ncards"))
	if err != nil {
		log.Warn("fail to read ncards")
	}
	templateStruct.NLands, err = CheckNLands(r.FormValue("nlands"))
	if err != nil {
		log.Warn("fail to read nlands")
	}
	// TODO check input validity
	// TODO deal with errors properly
	// must be > 0 and probably < 100
	templateStruct.White, err = strconv.Atoi(r.FormValue("white"))
	if err != nil {
		templateStruct.White = 0
	}
	templateStruct.Blue, err = strconv.Atoi(r.FormValue("blue"))
	if err != nil {
		templateStruct.Blue = 0
	}
	templateStruct.Black, err = strconv.Atoi(r.FormValue("black"))
	if err != nil {
		templateStruct.Black = 0
	}
	templateStruct.Red, err = strconv.Atoi(r.FormValue("red"))
	if err != nil {
		templateStruct.Red = 0
	}
	templateStruct.Green, err = strconv.Atoi(r.FormValue("green"))
	if err != nil {
		templateStruct.Green = 0
	}

	sumColored := templateStruct.White + templateStruct.Blue + templateStruct.Black + templateStruct.Red + templateStruct.Green
	if sumColored != 0 {
		var colors float64 = 0
		var minLandsPerColor float64

		// count how many colors we have
		var items []int = []int{templateStruct.White, templateStruct.Blue, templateStruct.Black, templateStruct.Red, templateStruct.Green}
		for _, value := range items {
			if value > 0 {
				colors += 1
			}
		}

		// if we have 4 or 5 colors, set "min lands per color" to 2
		// if we have 2 or 3 colors, set "min lands per color" to 3
		// if we only have 1 color, it doesn't matter
		if colors >= 4 {
			minLandsPerColor = 2
		} else {
			minLandsPerColor = 3
		}

		// compute the ratio of lands per color without lands we add to enforce minimums
		var ratio float64 = float64(sumColored) / (float64(templateStruct.NLands) - minLandsPerColor*colors)

		// for each color, add minLandsPerColor and ratio of remaining lands per color
		templateStruct.AWhite = computeLandsForColor(templateStruct.White, minLandsPerColor, ratio)
		templateStruct.ABlue = computeLandsForColor(templateStruct.Blue, minLandsPerColor, ratio)
		templateStruct.ABlack = computeLandsForColor(templateStruct.Black, minLandsPerColor, ratio)
		templateStruct.ARed = computeLandsForColor(templateStruct.Red, minLandsPerColor, ratio)
		templateStruct.AGreen = computeLandsForColor(templateStruct.Green, minLandsPerColor, ratio)
	}

	t, err := template.ParseFiles("templates/step2.html")
	if err != nil {
		log.Error(ErrTemplateParseFile)
	}

	err = t.Execute(w, templateStruct)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}
