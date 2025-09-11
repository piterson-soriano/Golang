package main

import "math"

//primeira letra maiúscula se torna pública (Vísivel dentro e fora do pacote)
//primeira letra minúscula se torna privado (Vísivel apenas dentro do pacote)

//Por exemplo
/*
Ponto - gerará algo público
ponto ou _Porto  - gerará algo privado
*/

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64 // Coordenada X (privada, pois começa com letra minúscula)
	y float64 // Coordenada Y (privada também)
}

// catetos calcula os catetos do triângulo formado entre dois pontos.
// Retorna a diferença absoluta entre as coordenadas X e Y.
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x) // Diferença absoluta entre os valores de X
	cy = math.Abs(b.y - a.y) // Diferença absoluta entre os valores de Y
	return                   // Retorna cx e cy

}

// Distancia calcula a distância entre dois pontos usando o Teorema de Pitágoras.
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)                             // Obtém os catetos entre os pontos
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2)) // Retorna a raiz quadrada da soma dos quadrados dos catetos

}
