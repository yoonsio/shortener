package shortener

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *App) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello")
}
