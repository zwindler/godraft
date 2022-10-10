package services

import (
	"errors"
	"math"
	"strconv"
)

var (
	ErrTemplateParseFile = errors.New("error while parsing template file")
	ErrTemplateExecute   = errors.New("error while execution of template file")
	ErrInvalidNNonLands  = errors.New("invalid NNonLands value")
	ErrInvalidNLands     = errors.New("invalid NLands value")
	ErrNoNNonLands       = errors.New("no value for NNonLands")
	ErrNoNLands          = errors.New("no value for NLands")
)

type TemplateStruct struct {
	DeckFormat  string
	DeckStyle   string
	NLands      int
	MinLands    int
	MaxLands    int
	NNonLands   int
	MinNonLands int
	MaxNonLands int
	MinCards    int
	MaxCards    int
	White       int
	AWhite      float64
	Blue        int
	ABlue       float64
	Black       int
	ABlack      float64
	Red         int
	ARed        float64
	Green       int
	AGreen      float64
	Version     string
}

func CheckNNonLands(NNonLands string) (NNonLandsInt int, err error) {
	if NNonLands == "" {
		return DefaultTemplateStruct.NNonLands, ErrNoNNonLands
	} else {
		NNonLandsInt, err = strconv.Atoi(NNonLands)
		if err != nil {
			return DefaultTemplateStruct.NNonLands, ErrInvalidNNonLands
		} else {
			if NNonLandsInt < DefaultTemplateStruct.MinCards || NNonLandsInt > DefaultTemplateStruct.MaxCards {
				return DefaultTemplateStruct.NNonLands, ErrInvalidNNonLands
			} else {
				return NNonLandsInt, nil
			}
		}
	}
}

func CheckNLands(NLands string) (NLandsInt int, err error) {
	if NLands == "" {
		return DefaultTemplateStruct.NLands, ErrNoNLands
	} else {
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
}

func (tmpl *TemplateStruct) SetDefaults(deckFormat string, deckStyle string) {
	tmpl.Version = Version

	if deckFormat == "standard" || deckFormat == "commander" {
		tmpl.DeckFormat = deckFormat
	} else {
		tmpl.DeckFormat = "draft"
	}

	if deckStyle == "aggro" || deckStyle == "control" {
		tmpl.DeckStyle = deckStyle
	} else {
		tmpl.DeckStyle = "midrange"
	}

	// define defaults depending on game type
	if tmpl.DeckFormat == "standard" {
		tmpl.MinCards = 60
		tmpl.MaxCards = 100
		tmpl.NLands = 22
	} else {
		if tmpl.DeckFormat == "commander" {
			tmpl.MinCards = 100
			tmpl.MaxCards = 100
			tmpl.NLands = 37
		} else {
			// draft
			tmpl.MinCards = 40
			tmpl.MaxCards = 60
			tmpl.NLands = 17
		}
	}

	// set sliders limits depending on game type
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
	tmpl.NNonLands = tmpl.MinCards - tmpl.NLands
	tmpl.MinNonLands = tmpl.MinCards - tmpl.MaxLands
	tmpl.MaxNonLands = tmpl.MaxCards - tmpl.MinLands
}

func computeLandsForColor(color int, minLandsPerColor float64, ratio float64) float64 {
	if color != 0 {
		return minLandsPerColor + math.Round(float64(color)/ratio*10)/10
	} else {
		return 0
	}
}

// TODO check input validity
// TODO deal with errors properly
// must be > 0 and probably < 100
func getLandFromForm(color string) int {
	value, err := strconv.Atoi(color)
	if err != nil {
		return 0
	}
	return value
}
