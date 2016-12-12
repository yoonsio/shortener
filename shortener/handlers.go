package shortener

import (
	"encoding/json"
	"log"
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
	log.Println(v)
	// TODO: push to db
	// TODO: fill cache
}

func (a *App) original(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// parse data
	var v OriginalReq
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		a.handleError(w, err)
		return
	}
	log.Println(v)
	key := v.Short
	// search from cache
	var b []byte
	err = a.cacheGroup.Get(nil, key, groupcache.AllocatingByteSliceSink(&b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// write result
	w.Write(b)
	w.Write([]byte{'\n'})
}
