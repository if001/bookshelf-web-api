// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"bookshelf-web-api/restapi/operations"
	"bookshelf-web-api/restapi/operations/book_description"
	"bookshelf-web-api/restapi/operations/books"
)

//go:generate swagger generate server --target ../../bookshelf-web-api --name BookshelfWeb --spec ../swagger.yaml

func configureFlags(api *operations.BookshelfWebAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BookshelfWebAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BooksDeleteBooksIDHandler == nil {
		api.BooksDeleteBooksIDHandler = books.DeleteBooksIDHandlerFunc(func(params books.DeleteBooksIDParams) middleware.Responder {
			return middleware.NotImplemented("operation books.DeleteBooksID has not yet been implemented")
		})
	}
	if api.BooksGetBooksHandler == nil {
		api.BooksGetBooksHandler = books.GetBooksHandlerFunc(func(params books.GetBooksParams) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooks has not yet been implemented")
		})
	}
	if api.BooksGetBooksIDHandler == nil {
		api.BooksGetBooksIDHandler = books.GetBooksIDHandlerFunc(func(params books.GetBooksIDParams) middleware.Responder {
			return middleware.NotImplemented("operation books.GetBooksID has not yet been implemented")
		})
	}
	if api.BookDescriptionGetBooksIDDescriptionHandler == nil {
		api.BookDescriptionGetBooksIDDescriptionHandler = book_description.GetBooksIDDescriptionHandlerFunc(func(params book_description.GetBooksIDDescriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation book_description.GetBooksIDDescription has not yet been implemented")
		})
	}
	if api.BooksPostBooksHandler == nil {
		api.BooksPostBooksHandler = books.PostBooksHandlerFunc(func(params books.PostBooksParams) middleware.Responder {
			return middleware.NotImplemented("operation books.PostBooks has not yet been implemented")
		})
	}
	if api.BookDescriptionPostBooksIDDescriptionHandler == nil {
		api.BookDescriptionPostBooksIDDescriptionHandler = book_description.PostBooksIDDescriptionHandlerFunc(func(params book_description.PostBooksIDDescriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation book_description.PostBooksIDDescription has not yet been implemented")
		})
	}
	if api.BooksPutBooksIDHandler == nil {
		api.BooksPutBooksIDHandler = books.PutBooksIDHandlerFunc(func(params books.PutBooksIDParams) middleware.Responder {
			return middleware.NotImplemented("operation books.PutBooksID has not yet been implemented")
		})
	}
	if api.BookDescriptionPutBooksIDDescriptionHandler == nil {
		api.BookDescriptionPutBooksIDDescriptionHandler = book_description.PutBooksIDDescriptionHandlerFunc(func(params book_description.PutBooksIDDescriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation book_description.PutBooksIDDescription has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
