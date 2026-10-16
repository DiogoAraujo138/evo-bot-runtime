package service_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EvolutionAPI/evo-bot-runtime/internal/empresa"
	"github.com/EvolutionAPI/evo-bot-runtime/pkg/dispatch/service"
	"github.com/EvolutionAPI/evo-bot-runtime/pkg/pipeline/model"
)

const empresaTeste = "11111111-1111-1111-1111-111111111111"

func TestDispatch_RepassaEmpresaAoCRM(t *testing.T) {
	var visto string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		visto = r.Header.Get(empresa.Header)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	eng := service.NewDispatchEngine("segredo")
	ctx := empresa.With(context.Background(), empresaTeste)
	if err := eng.Dispatch(ctx, 1, 1, "oi", model.BotConfig{}, server.URL); err != nil {
		t.Fatal(err)
	}
	if visto != empresaTeste {
		t.Fatalf("header %q, esperado %q", visto, empresaTeste)
	}

	visto = "x"
	if err := eng.Dispatch(context.Background(), 1, 1, "oi", model.BotConfig{}, server.URL); err != nil {
		t.Fatal(err)
	}
	if visto != "" { // Evo comunitário: nada é acrescentado
		t.Fatalf("sem empresa o header deveria faltar, veio %q", visto)
	}
}
