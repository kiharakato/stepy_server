package http

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/gorilla/sessions"
	"fmt"
)

const (
	get  = "GET"
	post = "POST"
)

func RequestGetParam(r *http.Request, key string) (string, bool) {
	value := r.URL.Query().Get(key)
	return value, (len(value) != 0)
}

func Chain(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return chain(true, true, true, f)
}

type customResponseWriter struct {
	io.Writer
	http.ResponseWriter
	status int
}

func (r *customResponseWriter) Write(b []byte) (int, error) {
	if r.Header().Get("Content-Type") == "" {
		r.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return r.Writer.Write(b)
}

func (r *customResponseWriter) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.status = status
}

func chain(log, cors, validate bool, f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ioWriter := w.(io.Writer)
		writer := &customResponseWriter{Writer: ioWriter, ResponseWriter: w, status: http.StatusOK}
		f(writer, r)
	})
}

type APIResource interface {
	Get(req *http.Request) (APIStatus, interface{})
	Post(req *http.Request) (APIStatus, interface{})
}

type APIStatus struct {
	success bool
	code    int
	message string
}

func APIResourceHandler(APIResource APIResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var status APIStatus
		var data interface{}

		switch r.Method {
		case get:
			status, data = APIResource.Get(r)
		case post:
			status, data = APIResource.Post(r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var content []byte
		var e error

		if !status.success {
			content, e = json.Marshal(apienvelope{
				Header: apiheader{Status: "fail", Message: status.message},
			})
		} else {
			content, e = json.Marshal(apienvelope{
				Header:   apiheader{Status: "success"},
				Response: data,
			})
		}
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)
		w.Write(content)
	}
}

type apiheader struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type apienvelope struct {
	Header   apiheader   `json:"header"`
	Response interface{} `json:"response"`
}

type APIResourceBase struct{}

func (APIResourceBase) Get(req *http.Request) (APIStatus, interface{}) {
	return APIStatus{success: false, code: http.StatusMethodNotAllowed, message: "server error"}, nil
}

func (APIResourceBase) Post(req *http.Request) (APIStatus, interface{}) {
	return APIStatus{success: false, code: http.StatusMethodNotAllowed, message: "server error"}, nil
}

func Success(code int) APIStatus {
	return APIStatus{success: true, code: code, message: ""}
}

func Fail(code int, message string) APIStatus {
	return APIStatus{success: false, code: code, message: message}
}

type Protocol struct {
	Wr  http.ResponseWriter
	Req *http.Request
	Session *sessions.Session
}

func (p Protocol) JsonWithInterface(data interface{}) {
	p.Wr.Header().Set("Content-Type", "application/json")

	content, e := json.Marshal(
		apienvelope{
			Header:   apiheader{Status: "success"},
			Response: data,
		})

	if e != nil {
		http.Error(p.Wr, e.Error(), http.StatusInternalServerError)
		return
	}

	p.Wr.Write(content)
}

func (p Protocol) Json(data []byte) {
	p.Wr.Header().Set("Content-Type", "application/json")

	content, e := json.Marshal(
		apienvelope{
			Header:   apiheader{Status: "success"},
			Response: data,
		})

	if e != nil {
		http.Error(p.Wr, e.Error(), http.StatusInternalServerError)
		return
	}

	p.Wr.Write(content)
}

func (p Protocol) SessionSave() {
	if err := p.Session.Save(p.Req, p.Wr); err != nil {
		fmt.Printf("Error saving session: %v", err)
		panic(err)
	}
}