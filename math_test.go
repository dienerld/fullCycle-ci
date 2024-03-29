package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(10, 20)

	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %f. Esperado: %f", total, 30.0)
	}
}
