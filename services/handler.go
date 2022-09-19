package services

import (
	"errors"
	"net/http"
)

var (
	ErrTemplateParseFile = errors.New("error while parsing template file")
	ErrTemplateExecute   = errors.New("error while execution of template file")
)

func Register() {
	http.HandleFunc("/", homeHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}
