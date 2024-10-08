package vhttp

import (
	"net/http"
	"strings"
)

type Mux struct {
	*http.ServeMux
}

func NewServeMux() *Mux {
	return &Mux{http.NewServeMux()}
}

func (m *Mux) Connect(pattern string, handler http.Handler) {
	m.handle(http.MethodConnect, pattern, handler)
}

func (m *Mux) Delete(pattern string, handler http.Handler) {
	m.handle(http.MethodDelete, pattern, handler)
}

func (m *Mux) Get(pattern string, handler http.Handler) {
	m.handle(http.MethodGet, pattern, handler)
}

func (m *Mux) Head(pattern string, handler http.Handler) {
	m.handle(http.MethodHead, pattern, handler)
}

func (m *Mux) Options(pattern string, handler http.Handler) {
	m.handle(http.MethodOptions, pattern, handler)
}

func (m *Mux) Post(pattern string, handler http.Handler) {
	m.handle(http.MethodPost, pattern, handler)
}

func (m *Mux) Put(pattern string, handler http.Handler) {
	m.handle(http.MethodPut, pattern, handler)
}

func (m *Mux) Patch(pattern string, handler http.Handler) {
	m.handle(http.MethodPatch, pattern, handler)
}

func (m *Mux) Trace(pattern string, handler http.Handler) {
	m.handle(http.MethodTrace, pattern, handler)
}

func (m *Mux) handle(method string, pattern string, handler http.Handler) {
	m.Handle(strings.Join([]string{method, pattern}, " "), handler)
}

// HandlerFunc

func (m *Mux) FuncConnect(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodConnect, pattern, handler)
}

func (m *Mux) FuncDelete(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodDelete, pattern, handler)
}

func (m *Mux) FuncGet(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodGet, pattern, handler)
}

func (m *Mux) FuncHead(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodHead, pattern, handler)
}

func (m *Mux) FuncOptions(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodOptions, pattern, handler)
}

func (m *Mux) FuncPost(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodPost, pattern, handler)
}

func (m *Mux) FuncPut(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodPut, pattern, handler)
}

func (m *Mux) FuncPatch(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodPatch, pattern, handler)
}

func (m *Mux) FuncTrace(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.handleFunc(http.MethodTrace, pattern, handler)
}

func (m *Mux) handleFunc(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.HandleFunc(strings.Join([]string{method, pattern}, " "), handler)
}
