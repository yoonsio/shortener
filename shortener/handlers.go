package shortener

import (
	"encoding/json"
	"net/http"

	"github.com/golang/groupcache"
	"github.com/julienschmidt/httprouter"
)

func (a *App) shorten(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// parse data
	var v ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		a.handleError(w, err)
		return
	}

	// TODO: validate ShortenRequest

	// generate short URI (random?)
	// TODO: get current URL from configuration
	// TODO: optimize string cocatenation
	short := "https://localhost:9000/" + GenerateRandomString(8)

	// push to db
	err = a.db.Register(v.URL, short)
	if err != nil {
		a.handleError(w, err)
		return
	}

	// return correct response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(ShortenResponse{short})
	if err != nil {
		a.handleError(w, err)
		return
	}
}

func (a *App) original(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// parse data
	var v OriginalRequest
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		a.handleError(w, err)
		return
	}

	// TODO: validate OriginalReq

	// search from cache (which in turn search database)
	var b []byte
	err = a.cacheGroup.Get(nil, v.Short, groupcache.AllocatingByteSliceSink(&b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// return correct response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(OriginalResponse{string(b)})
	if err != nil {
		a.handleError(w, err)
		return
	}

}
