package main

import "testing"

func TestSoma(t *testing.T) {

	somatoria := soma(5, 5) //somatória deveria retorna 10

	if somatoria != 10 {
		t.Errorf("soma (5,5) failed, expected %v, got %v", 10, somatoria)
	}

}
