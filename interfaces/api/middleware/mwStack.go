package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MwStack struct {
	middlewares []Middleware
}

func NewMws(mws ...Middleware) MwStack {
	return MwStack{append([]Middleware(nil), mws...)}
}

func (m MwStack) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i](h)
	}
	return h
}