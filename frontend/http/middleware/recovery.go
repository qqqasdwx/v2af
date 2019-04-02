package middleware

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/qqqasdwx/v2af/frontend/http/render"
	"github.com/toolkits/web/errors"
)

// Recovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.
type Recovery struct {
	PrintStack bool
	StackAll   bool
	StackSize  int
}

// NewRecovery returns a new instance of Recovery
func NewRecovery() *Recovery {
	return &Recovery{
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
	}
}

func (rec *Recovery) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {

			if customError, ok := err.(errors.Error); ok {

				if isAjax(r) || strings.HasPrefix(r.RequestURI, "/api") {
					render.JSON(w, map[string]interface{}{"msg": customError.Msg})
					return
				}

				if customError.Code == http.StatusUnauthorized || customError.Code == http.StatusForbidden {
					http.Redirect(w, r, "/auth/signin?callback="+r.URL.String(), 302)
					return
				}

				render.Data(r, "Error", customError.Msg)
				render.HTML(r, w, "common/error")
				return
			}

			// Negroni part
			w.WriteHeader(http.StatusInternalServerError)
			stack := make([]byte, rec.StackSize)
			stack = stack[:runtime.Stack(stack, rec.StackAll)]

			f := "PANIC: %s\n%s"
			log.Printf(f, err, stack)

			if rec.PrintStack {
				fmt.Fprintf(w, f, err, stack)
			}
		}
	}()

	next(w, r)
}

func isAjax(r *http.Request) bool {
	return r.Header.Get("X-Requested-With") == "XMLHttpRequest"
}
