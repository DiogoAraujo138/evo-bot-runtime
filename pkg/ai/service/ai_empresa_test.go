package service_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EvolutionAPI/evo-bot-runtime/internal/empresa"
	aiModel "github.com/EvolutionAPI/evo-bot-runtime/pkg/ai/model"
	aiService "github.com/EvolutionAPI/evo-bot-runtime/pkg/ai/service"
)

func TestCall_RepassaEmpresaAoProcessor(t *testing.T) {
	const id = "11111111-1111-1111-1111-111111111111"
	var visto []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		visto = append(visto, r.Header.Get(empresa.Header))
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":"1","result":{}}`))
	}))
	defer server.Close()

	adapter := aiService.NewAIAdapter(5, 0, 1)
	req := &aiModel.A2ARequest{OutgoingURL: server.URL + "/api/v1/a2a/agent-1", ContactID: 1, ConversationID: 1, ApiKey: "k", Message: "oi"}
	_, _ = adapter.Call(empresa.With(context.Background(), id), req) // a resposta mínima pode não normalizar: só o cabeçalho importa
	_, _ = adapter.Call(context.Background(), req)

	if len(visto) != 2 || visto[0] != id || visto[1] != "" {
		t.Fatalf("cabeçalhos vistos: %q", visto)
	}
}
