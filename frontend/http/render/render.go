package render

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/qqqasdwx/v2af/frontend/config"
	// "github.com/qqqasdwx/v2af/frontend/http/cookie"

	"github.com/gorilla/context"
	"github.com/unrolled/render"
)

var Render *render.Render

func Init() {
	debug := config.Config().Debug
	Render = render.New(render.Options{
		Directory:     "views",
		Extensions:    []string{".html"},
		Delims:        render.Delims{"{{", "}}"},
		IndentJSON:    false,
		IsDevelopment: debug,
	})
}

func Data(r *http.Request, key string, val interface{}) {
	m, ok := context.GetOk(r, "DATA_MAP")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		context.Set(r, "DATA_MAP", mm)
	} else {
		context.Set(r, "DATA_MAP", map[string]interface{}{key: val})
	}
}

func HTML(r *http.Request, w http.ResponseWriter, name string, htmlOpt ...render.HTMLOptions) {
	// userid, username, found := cookie.ReadUser(r)
	// Data(r, "ID", userid)
	// Data(r, "NAME", username)
	// Data(r, "LOGIN", found)
	Render.HTML(w, http.StatusOK, name, context.Get(r, "DATA_MAP"), htmlOpt...)
}

func JSON(w http.ResponseWriter, v interface{}, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	// Render.JSON(w, code, v)

	bs, _ := json.Marshal(v)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func Auto(w http.ResponseWriter, err error, v ...interface{}) {
	if err != nil {
		JSON(w, map[string]interface{}{"msg": err.Error()})
		return
	}

	if len(v) > 0 {
		JSON(w, map[string]interface{}{"msg": "", "data": v[0]})
	} else {
		JSON(w, map[string]interface{}{"msg": ""})
	}
}

func Text(w http.ResponseWriter, v string, codes ...int) {
	code := http.StatusOK
	if len(codes) > 0 {
		code = codes[0]
	}
	Render.Text(w, code, v)
}
