package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/ruhollahh/go-lazy-web-components/api/handler"
	"github.com/ruhollahh/go-lazy-web-components/api/middleware"
)

func (a *API) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./web/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	handler := handler.New(a.Logger)
	middleware := middleware.New(a.Logger)

	router.HandlerFunc(http.MethodGet, "/", handler.Home)
	router.HandlerFunc(http.MethodGet, "/show-random-lazy-comp", handler.ShowRandomLazyComp)

	standard := alice.New(middleware.RecoverPanic, middleware.LogRequest, middleware.SecureHeaders)
	return standard.Then(router)
}
