package services

import (
	"math"
	"strconv"
)

type TemplateStruct struct {
	DeckFormat string
	DeckStyle  string
	NCards     int
	MinCards   int
	MaxCards   int
	NLands     int
	MinLands   int
	MaxLands   int
	White      int
	AWhite     float64
	Blue       int
	ABlue      float64
	Black      int
	ABlack     float64
	Red        int
	ARed       float64
	Green      int
	AGreen     float64
	Version    string
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

func (tmpl *TemplateStruct) SetDefaults(deckFormat string, deckStyle string) {
	tmpl.Version = Version

	if deckFormat == "standard" || deckFormat == "commander" {
		tmpl.DeckFormat = deckFormat
	} else {
		tmpl.DeckFormat = "draft"
	}

	if deckStyle == "aggro" || deckStyle == "control" {
		tmpl.DeckFormat = deckFormat
	} else {
		tmpl.DeckFormat = "midrange"
	}

	// define defaults depending on game type
	if tmpl.DeckFormat == "standard" {
		tmpl.NCards = 60
		tmpl.MinCards = 60
		tmpl.MaxCards = 100
		tmpl.NLands = 22
	} else {
		if tmpl.DeckFormat == "commander" {
			tmpl.NCards = 100
			tmpl.MinCards = 100
			tmpl.MaxCards = 100
			tmpl.NLands = 37
		} else {
			// draft
			tmpl.NCards = 40
			tmpl.MinCards = 40
			tmpl.MaxCards = 60
			tmpl.NLands = 17
		}
	}

	// set limits depending on game type
	tmpl.MinLands = int(float64(tmpl.NLands) * 0.5)
	tmpl.MaxLands = int(float64(tmpl.NLands) * 1.5)

	// adapt defaults depending on game style
	if tmpl.DeckStyle == "aggro" {
		tmpl.NLands = int(float64(tmpl.NLands) * 0.9)
	} else {
		if tmpl.DeckStyle == "control" {
			tmpl.NLands = int(float64(tmpl.NLands) * 1.1)
		}
	}
}

func computeLandsForColor(color int, minLandsPerColor float64, ratio float64) float64 {
	if color != 0 {
		return minLandsPerColor + math.Round(float64(color)/ratio*10)/10
	} else {
		return 0
	}
}
