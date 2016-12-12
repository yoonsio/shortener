package shortener

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/burntsushi/toml"
	"github.com/golang/groupcache"
	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

// App is main web application
type App struct {
	*httprouter.Router                   // httprouter
	handlers           http.Handler      // modified net/http handler
	config             Config            // configuration struct
	db                 *MongoClient      // database client
	cacheGroup         *groupcache.Group // groupcache group
}

// NewApp creates new ShortenerApp
func NewApp(configFile string) *App {

	// initialize empty ShortenerApp struct
	app := App{
		Router: httprouter.New(),
	}

	// load config file if exists
	if configFile != "" {
		if _, err := toml.DecodeFile(configFile, &app.config); err != nil {
			log.Panic(err)
		}
	}

	// establish db connection
	app.db = NewMongoClient(app.config.Database.URI, app.config.Database.DBName)

	// configure groupcache
	app.cacheGroup = groupcache.NewGroup(
		"shortener",
		64<<20,
		groupcache.GetterFunc(
			func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
				// TODO: look for value from database
				var v []byte
				ok := false
				// if not found, return error
				if !ok {
					return errors.New("failed to find a key")
				}
				dest.SetBytes(v)
				return nil
			},
		),
	)

	// configure groupcache pool
	cachePool := groupcache.NewHTTPPool("127.0.0.1")
	// cachePool.Set(p...)
	_ = cachePool

	// TODO: add handlers
	app.POST("/shorten", app.shorten)
	app.GET("/original", app.original)

	// add static resources handler
	fs := http.FileServer(http.Dir("static"))
	app.Handler("GET", "/static/", http.StripPrefix("/static/", fs))

	// add middlewares
	h := handlers.LoggingHandler(os.Stdout, app)
	h = handlers.ProxyHeaders(h)
	h = handlers.CompressHandler(h)
	h = handlers.RecoveryHandler()(h)
	app.handlers = h
	return &app
}

// Handle handler errors
func (a *App) handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// Run runs this web application
func (a *App) Run() {
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", a.config.Server.Port), a.handlers),
	)
}
