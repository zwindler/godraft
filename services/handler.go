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
		templateStruct.NNonLands, err = CheckNNonLands(nnonlands)
		if err != nil {
			log.Warn("fail to read nnonlands")
		}
	}
	if nlands != "" {
		templateStruct.NLands, err = CheckNLands(nlands)
		if err != nil {
			log.Warn("fail to read nlands")
		}
	}

	templateStruct.White = getLandFromForm(r.FormValue("white"))
	templateStruct.Blue = getLandFromForm(r.FormValue("blue"))
	templateStruct.Black = getLandFromForm(r.FormValue("black"))
	templateStruct.Red = getLandFromForm(r.FormValue("red"))
	templateStruct.Green = getLandFromForm(r.FormValue("green"))

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
