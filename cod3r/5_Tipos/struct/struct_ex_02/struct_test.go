package main

import (
	"testing"
)

func TestPrecoComDesconto(t *testing.T) {
	c := carro{
		marca:  "Toyota",
		modelo: "Corolla Cross 2025",
		preco:  164000.00,
	}
	want := 164000.00 * (1 - 0.18)
	got := c.precoComDesconto()
	if got != want {
		t.Errorf("precoComDesconto() = %v; want %v", got, want)
	}
}

func TestCarroFields(t *testing.T) {
	c := carro{
		marca:  "Toyota",
		modelo: "Corolla Cross 2025",
		preco:  164000.00,
	}
	if c.marca != "Toyota" {
		t.Errorf("marca = %v; want %v", c.marca, "Toyota")
	}
	if c.modelo != "Corolla Cross 2025" {
		t.Errorf("modelo = %v; want %v", c.modelo, "Corolla Cross 2025")
	}
	if c.preco != 164000.00 {
		t.Errorf("preco = %v; want %v", c.preco, 164000.00)
	}
}
