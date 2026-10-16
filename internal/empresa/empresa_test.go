package empresa

import (
	"context"
	"testing"
)

func TestWithFrom(t *testing.T) {
	base := context.Background()
	if From(base) != "" {
		t.Fatal("contexto vazio não tem empresa")
	}
	if With(base, "") != base {
		t.Fatal("id vazio não deve alterar o contexto")
	}
	if got := From(With(base, "11111111-1111-1111-1111-111111111111")); got != "11111111-1111-1111-1111-111111111111" {
		t.Fatalf("got %q", got)
	}
}
