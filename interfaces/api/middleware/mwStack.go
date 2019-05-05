package middleware

import (
	"github.com/julienschmidt/httprouter"
)

type Middleware func(handler httprouter.Handle) httprouter.Handle

type MwStack struct {
	middlewares []Middleware
}

func NewMws(mws ...Middleware) MwStack {
	return MwStack{append([]Middleware(nil), mws...)}
}

func (m MwStack) Then(h httprouter.Handle) httprouter.Handle {
	for i := range m.middlewares {
		h = m.middlewares[len(m.middlewares)-1-i](h)
	}
	return h
}

