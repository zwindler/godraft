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
	Version string
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
		var ratio float64 = float64(sumColored) / float64(templateStruct.NLands)

		templateStruct.AWhite = math.Round(float64(templateStruct.White)/ratio*10) / 10
		templateStruct.ABlue = math.Round(float64(templateStruct.Blue)/ratio*10) / 10
		templateStruct.ABlack = math.Round(float64(templateStruct.Black)/ratio*10) / 10
		templateStruct.ARed = math.Round(float64(templateStruct.Red)/ratio*10) / 10
		templateStruct.AGreen = math.Round(float64(templateStruct.Green)/ratio*10) / 10
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
		return 40, ErrInvalidNCards
	} else {
		if NCardsInt < DefaultTemplateStruct.MinCards || NCardsInt > DefaultTemplateStruct.MaxCards {
			return 40, ErrInvalidNCards
		} else {
			return NCardsInt, nil
		}
	}
}

func CheckNLands(NLands string) (NLandsInt int, err error) {
	NLandsInt, err = strconv.Atoi(NLands)
	if err != nil {
		return 17, ErrInvalidNLands
	} else {
		if NLandsInt < DefaultTemplateStruct.MinLands || NLandsInt > DefaultTemplateStruct.MaxLands {
			return 17, ErrInvalidNLands
		} else {
			return NLandsInt, nil
		}
	}
}

func (tmpl *TemplateStruct) SetDefaults() {
	tmpl.Version = Version
	tmpl.NCards = 40
	tmpl.NLands = 17
	tmpl.White = 0
	tmpl.Blue = 0
	tmpl.Black = 0
	tmpl.Red = 0
	tmpl.Green = 0
	tmpl.MinCards = 40
	tmpl.MaxCards = 60
	tmpl.MinLands = 13
	tmpl.MaxLands = 25
}
