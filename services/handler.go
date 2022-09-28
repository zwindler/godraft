package services

import (
	"errors"
	"html/template"
	"math"
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

type TemplateStruct struct {
	NCards   int
	NLands   int
	White    int
	Blue     int
	Black    int
	Red      int
	Green    int
	MinCards int
	MaxCards int
	MinLands int
	MaxLands int
	AWhite   float64
	ABlue    float64
	ABlack   float64
	ARed     float64
	AGreen   float64
	Version  string
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

	err = t.Execute(w, DefaultTemplateStruct)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}

func computeHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var templateStruct TemplateStruct
	templateStruct.SetDefaults()

	// TODO deal with errors
	templateStruct.NCards, err = CheckNCards(r.FormValue("ncards"))
	if err != nil {
		return
	}
	templateStruct.NLands, err = CheckNLands(r.FormValue("nlands"))
	if err != nil {
		return
	}
	// TODO check input validity
	// must be > 0 and probably < 100
	templateStruct.White, err = strconv.Atoi(r.FormValue("white"))
	if err != nil {
		return
	}
	templateStruct.Blue, err = strconv.Atoi(r.FormValue("blue"))
	if err != nil {
		return
	}
	templateStruct.Black, err = strconv.Atoi(r.FormValue("black"))
	if err != nil {
		return
	}
	templateStruct.Red, err = strconv.Atoi(r.FormValue("red"))
	if err != nil {
		return
	}
	templateStruct.Green, err = strconv.Atoi(r.FormValue("green"))
	if err != nil {
		return
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

		// computate the ratio of lands per color without lands we add to enforce minimums
		var ratio float64 = float64(sumColored) / (float64(templateStruct.NLands) - minLandsPerColor * colors)

		// for each color, add minLandsPerColor and ratio of remaining lands per color
		templateStruct.AWhite = computateLandsForColor(templateStruct.White, minLandsPerColor, ratio)
		templateStruct.ABlue = computateLandsForColor(templateStruct.Blue, minLandsPerColor, ratio)
		templateStruct.ABlack = computateLandsForColor(templateStruct.Black, minLandsPerColor, ratio)
		templateStruct.ARed = computateLandsForColor(templateStruct.Red, minLandsPerColor, ratio)
		templateStruct.AGreen = computateLandsForColor(templateStruct.Green, minLandsPerColor, ratio)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(ErrTemplateParseFile)
	}

	err = t.Execute(w, templateStruct)
	if err != nil {
		log.Error(ErrTemplateExecute)
	}
}

func CheckNCards(NCards string) (NCardsInt int, err error) {
	NCardsInt, err = strconv.Atoi(NCards)
	if err != nil {
		return DefaultTemplateStruct.NCards, ErrInvalidNCards
	} else {
		if NCardsInt < DefaultTemplateStruct.MinCards || NCardsInt > DefaultTemplateStruct.MaxCards {
			return DefaultTemplateStruct.NCards, ErrInvalidNCards
		} else {
			return NCardsInt, nil
		}
	}
}

func CheckNLands(NLands string) (NLandsInt int, err error) {
	NLandsInt, err = strconv.Atoi(NLands)
	if err != nil {
		return DefaultTemplateStruct.NLands, ErrInvalidNLands
	} else {
		if NLandsInt < DefaultTemplateStruct.MinLands || NLandsInt > DefaultTemplateStruct.MaxLands {
			return DefaultTemplateStruct.NLands, ErrInvalidNLands
		} else {
			return NLandsInt, nil
		}
	}
}

func (tmpl *TemplateStruct) SetDefaults() {
	tmpl.Version = Version
	tmpl.MinCards = 20
	tmpl.NCards = 23
	tmpl.MaxCards = 40
	tmpl.MinLands = 13
	tmpl.NLands = 17
	tmpl.MaxLands = 25
	tmpl.White = 0
	tmpl.Blue = 0
	tmpl.Black = 0
	tmpl.Red = 0
	tmpl.Green = 0
}

func computateLandsForColor(color int, minLandsPerColor float64, ratio float64) (float64) {
	if color != 0 {
		return minLandsPerColor + math.Round(float64(color)/ratio*10) / 10
	} else {
		return 0
	}
}