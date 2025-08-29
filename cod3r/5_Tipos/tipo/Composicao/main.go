package main

import "fmt"

//Interface "esportivo" com um único método
type esportivo interface {
	ligarTurbo()
}

//Interface "luxuoso" com um único método
type luxuoso interface {
	fazerBaliza()
}

//Esta é um interface composta "esportivoLuxuoso"
//Incorpora as interfaces esportivo e luxuoso
//Qualquer tipo que implementar ambos os métodos será compatível com essa interface
type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//Pode adicionar mais métodos
}

//Define um tipo chamado bwm7 como uma struct vazia.
//Esse tipo será usado para implementar os métodos das interfaces.
type bwm7 struct{}

//Implementa o método "ligarTurbo" para o tipo bwm7.
//Isso faz com que "bwm7" seja compatível com a interface "esportivo".
func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b esportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
