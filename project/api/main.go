package main

import "fmt"

func main() {

}

func Test() {
	var a int = 5
	var b int = 6
	// var c int = a + b

	c := Suma(a, b)
	d := Resta(a, b)
	e, f := SumaResta(a, b)
	g := Potencia(a, b)

	fmt.Println("Hello, World!")
	Format("La suma de %d + %d es %d", a, b, c)
	Format("La resta de %d - %d es %d", a, b, d)
	Format("La suma de %d + %d es %d y la resta de %d - %d es %d", a, b, e, a, b, f)
	Format("La potencia de %d elevado a %d es %d", a, b, g)
}

func Suma(a int, b int) int {
	return a + b
}

func Resta(a int, b int) int {
	return a - b
}
func SumaResta(a int, b int) (int, int) {
	c := Suma(a, b)
	d := Resta(a, b)

	return c, d
}

func Potencia(a int, b int) int {
	for i := 1; i < b; i++ {
		a = a * a
	}

	return a
}

func Format(formato string, args ...any) {
	result := fmt.Sprintf(formato, args...)
	fmt.Println(result)
}
