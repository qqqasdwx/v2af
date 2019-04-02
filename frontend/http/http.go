package http

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/qqqasdwx/v2af/frontend/config"
	// "github.com/qqqasdwx/v2af/frontend/http/cookie"
	"github.com/qqqasdwx/v2af/frontend/http/middleware"
	// "github.com/qqqasdwx/v2af/frontend/http/render"
	"github.com/v2af/file"
)

func Start() {
	// render.Init()
	// cookie.Init()

	r := mux.NewRouter().StrictSlash(false)
	ConfigRouter(r)

	n := negroni.New()
	n.Use(middleware.NewRecovery())
	n.Use(middleware.NewLogger(file.MustOpenLogFile(config.Config().Http.Access)))

	n.UseHandler(r)
	n.Run(config.Config().Http.Listen)
}
