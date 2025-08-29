package main

import "fmt"

type temperaturaCelsius float64

func (t temperaturaCelsius) ParaFahrenheit() float64 {
	return float64(t)*1.8 + 32
}

func main() {
	t := temperaturaCelsius(30)
	fmt.Println("Temperatura em Fahrenheit", t.ParaFahrenheit())
}
