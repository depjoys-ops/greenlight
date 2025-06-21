package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movies", app.requirePermission("movies:read", app.listMoviesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requirePermission("movies:read", app.showMovieHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requirePermission("movies:write", app.updateMovieHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requirePermission("movies:write", app.deleteMovieHandler))

	// $ BODY='{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'
	// $ curl -i -d "$BODY" localhost:4000/v1/users
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	// $ curl -X PUT -d '{"token": "invalid"}' localhost:4000/v1/users/activated
	// $ curl -X PUT -d '{"token": "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}' localhost:4000/v1/users/activated
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	// $ curl -d '{"email": "alice@example.com", "password": "pa55word"}' localhost:4000/v1/tokens/authentication
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	// $ curl -H "Authorization: Bearer ZIRXR4SAAUY6DOISEOTUI4L7CS" localhost:4000/v1/healthcheck
	// $ curl -i -H "Authorization: INVALID" localhost:4000/v1/healthcheck

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
