package main

type Car struct {
	Model string
	Color string
}

//metodo da Struct
func (c Car) Start() {
	println(c.Model + " has been started")
}

//Porém as alterações que estão sendo feitas é uma cópia que será descartada logo após o uso.
func (c *Car) changeColor(color string) { //Quando usamos o ponteiro indicamos que a variável não deve ser uma copia, mas sim a ref da Struct
	c.Color = color
	println("New color: " + c.Color)
}

func main() {

}
