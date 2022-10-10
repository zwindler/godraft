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

func (tmpl *TemplateStruct) CheckNNonLands(NNonLands string) (err error) {
	if NNonLands == "" {
		// if empty, return and keep defaults
		return ErrNoNNonLands
	} else {
		NNonLandsInt, err := strconv.Atoi(NNonLands)
		if err != nil {
			// str to int failure
			return ErrInvalidNNonLands
		} else {
			if NNonLandsInt < tmpl.MinNonLands || NNonLandsInt > tmpl.MaxCards {
				// nnonlands doesn't fit in format boundaries
				return ErrInvalidNNonLands
			} else {
				tmpl.NNonLands = NNonLandsInt
				return nil
			}
		}
	}
}

func (tmpl *TemplateStruct) CheckNLands(NLands string) (err error) {
	if NLands == "" {
		// if empty, return and keep defaults
		return ErrNoNLands
	} else {
		NLandsInt, err := strconv.Atoi(NLands)
		if err != nil {
			// str to int failure
			return ErrInvalidNLands
		} else {
			if NLandsInt < tmpl.MinLands || NLandsInt > tmpl.MaxLands {
				// nlands doesn't fit in format boundaries
				return ErrInvalidNLands
			} else {
				tmpl.NLands = NLandsInt
				return nil
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
		tmpl.MaxCards = 80
		tmpl.NLands = 22
	} else {
		if tmpl.DeckFormat == "commander" {
			tmpl.MinCards = 100
			tmpl.MaxCards = 100
			tmpl.NLands = 37
		} else {
			// draft
			tmpl.MinCards = 40
			tmpl.MaxCards = 52
			tmpl.NLands = 17
		}
	}

	// set sliders limits depending on game type
	tmpl.MinLands = int(float64(tmpl.NLands) * 0.75)
	tmpl.MaxLands = int(float64(tmpl.NLands) * 1.33)

	// adapt defaults depending on game style
	if tmpl.DeckStyle == "aggro" {
		tmpl.NLands = int(float64(tmpl.NLands) * 0.9)
	} else {
		if tmpl.DeckStyle == "control" {
			tmpl.NLands = int(float64(tmpl.NLands) * 1.1)
		}
	}

	// set boundaries for lands
	tmpl.NNonLands = tmpl.MinCards - tmpl.NLands
	tmpl.MinNonLands = tmpl.MinCards - tmpl.MaxLands
	// commander is more restrictive so I must set different rules
	if deckFormat == "commander" {
		tmpl.MaxNonLands = tmpl.MaxCards - tmpl.MinLands
	} else {
		// I artificially allow more cards by allowing
		tmpl.MaxNonLands = tmpl.MaxCards - tmpl.MaxLands
	}

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
func getLandsFromForm(color string) int {
	value, err := strconv.Atoi(color)
	if err != nil {
		return 0
	}
	return value
}

func (tmpl *TemplateStruct) suggestLands() {
	sumColored := tmpl.White + tmpl.Blue + tmpl.Black + tmpl.Red + tmpl.Green
	if sumColored != 0 {
		var colors float64 = 0
		var minLandsPerColor float64

		// count how many colors we have
		var items []int = []int{tmpl.White, tmpl.Blue, tmpl.Black, tmpl.Red, tmpl.Green}
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
		var ratio float64 = float64(sumColored) / (float64(tmpl.NLands) - minLandsPerColor*colors)

		// for each color, add minLandsPerColor and ratio of remaining lands per color
		tmpl.AWhite = computeLandsForColor(tmpl.White, minLandsPerColor, ratio)
		tmpl.ABlue = computeLandsForColor(tmpl.Blue, minLandsPerColor, ratio)
		tmpl.ABlack = computeLandsForColor(tmpl.Black, minLandsPerColor, ratio)
		tmpl.ARed = computeLandsForColor(tmpl.Red, minLandsPerColor, ratio)
		tmpl.AGreen = computeLandsForColor(tmpl.Green, minLandsPerColor, ratio)
	}
}
