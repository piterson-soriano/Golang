package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {

	/*
		O comando  em Go (Golang) é usado para verificar, em tempo de execução,
		se o programa está sendo executado em uma arquitetura de 64 bits baseada em AMD/Intel (x86-64).
	*/
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	//...
	t.Fail()
}
