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
	ErrTemplateParseFile = errors.New("error while parsing template file")
	ErrTemplateExecute   = errors.New("error while execution of template file")
	ErrInvalidNCards     = errors.New("invalid ncards value")
	ErrInvalidNLands     = errors.New("invalid nlands value")
	EmptyCompute         = Compute{
		NCards: 40,
		NLands: 17,
		White:  0,
		Blue:   0,
		Black:  0,
		Red:    0,
		Green:  0,
	}
	MinCards = 40
	MaxCards = 60
	MinLands = 13
	MaxLands = 25
)

type Compute struct {
	NCards int
	NLands int
	White  int
	Blue   int
	Black  int
	Red    int
	Green  int
	AWhite float64
	ABlue  float64
	ABlack float64
	ARed   float64
	AGreen float64
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

func CheckNCards(ncards string) (ncardsi int, err error) {
	ncardsi, err = strconv.Atoi(ncards)
	if err != nil {
		return 40, ErrInvalidNCards
	} else {
		if ncardsi < MinCards || ncardsi > MaxCards {
			return 40, ErrInvalidNCards
		} else {
			return ncardsi, nil
		}
	}
}

func CheckNLands(nlands string) (nlandsi int, err error) {
	nlandsi, err = strconv.Atoi(nlands)
	if err != nil {
		return 17, ErrInvalidNLands
	} else {
		if nlandsi < MinLands || nlandsi > MaxLands {
			return 17, ErrInvalidNLands
		} else {
			return nlandsi, nil
		}
	}
}
