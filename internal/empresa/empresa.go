// Package empresa carrega a empresa (multi-empresa Chordata) pelo contexto até as chamadas de saída.
// O CRM manda `empresa_id` no evento; sem ele (Evo comunitário) nada muda.
package empresa

import "context"

// Header é o mesmo cabeçalho usado pelo CRM, Flow, Core e Processor.
const Header = "X-Evo-Tenant-Id"

type chave struct{}

// With devolve um contexto que carrega a empresa (id vazio devolve o contexto como veio).
func With(ctx context.Context, id string) context.Context {
	if id == "" {
		return ctx
	}
	return context.WithValue(ctx, chave{}, id)
}

// From devolve a empresa do contexto, ou "".
func From(ctx context.Context) string {
	id, _ := ctx.Value(chave{}).(string)
	return id
}
