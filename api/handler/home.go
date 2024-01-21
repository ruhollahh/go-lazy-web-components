package handler

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/ruhollahh/go-lazy-web-components/web/views/components"
	"github.com/ruhollahh/go-lazy-web-components/web/views/pages"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(context.Background(), w)
}

func (h *Handler) ShowRandomLazyComp(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(2)
	if random == 0 {
		components.LazyCompOne().Render(context.Background(), w)
	} else {
		components.LazyCompTwo().Render(context.Background(), w)
	}
}
