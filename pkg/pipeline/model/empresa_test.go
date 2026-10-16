package model_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/EvolutionAPI/evo-bot-runtime/pkg/pipeline/model"
)

func TestEmpresaNoEventoENoEstado(t *testing.T) {
	var ev model.MessageEvent
	if err := json.Unmarshal([]byte(`{"contact_id":1,"conversation_id":2,"postback_url":"http://crm/x","empresa_id":"abc"}`), &ev); err != nil {
		t.Fatal(err)
	}
	if ev.EmpresaID != "abc" {
		t.Fatalf("empresa_id não lido: %q", ev.EmpresaID)
	}
	// o estado guardado no Redis precisa levar a empresa (recuperação após reinício)
	b, _ := json.Marshal(model.PipelineState{EmpresaID: "abc"})
	var volta model.PipelineState
	_ = json.Unmarshal(b, &volta)
	if volta.EmpresaID != "abc" {
		t.Fatalf("estado perdeu a empresa: %s", b)
	}
	// Evo comunitário: sem empresa o campo nem aparece
	b, _ = json.Marshal(model.PipelineState{})
	if strings.Contains(string(b), "empresa_id") {
		t.Fatalf("estado sem empresa mudou de formato: %s", b)
	}
}
