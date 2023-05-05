package main

import "fmt"

type objeto struct {
	numero int
	nombre string
	precio int
}

type objetoLista struct {
	objetos []*objeto
}

type objetoAgrega struct {
	objetosListaNueva []*objeto
}

// func (t *objetoLista) agregarALista(tl *objeto) {
// 	t.objetos = append(t.objetos, tl)
// }

func (t *objetoLista) imprimirLista() {
	fmt.Println("Los Productos de la lista son los siguientes: ")
	for _, producto := range t.objetos {
		fmt.Println("#", producto.numero, "Nombre", producto.nombre, "Precio", producto.precio)
	}
	fmt.Println("Por favor seleccione el producto a ingresar: ")
}

func (t *objetoLista) agregarObjetos() {
	var eleccion int
	fmt.Scanln(&eleccion)
	//eleccion := 2
	for _, producto := range t.objetos {
		if producto.numero == eleccion {
			tn := &objeto{
				producto.numero,
				producto.nombre,
				producto.precio,
			}
			agregar := &objetoAgrega{}
			agregar.agregaobjetosLista(tn)
			agregar.imprimirListaAgregados()
		}
	}

}

func (t *objetoAgrega) agregaobjetosLista(tl *objeto) {
	t.objetosListaNueva = append(t.objetosListaNueva, tl)

}

func (t *objetoAgrega) imprimirListaAgregados() {
	fmt.Println("Los Productos agregados a su lista son los siguientes: ")
	for _, productoNue := range t.objetosListaNueva {
		fmt.Println("#", productoNue.numero, "Nombre", productoNue.nombre, "Precio", productoNue.precio)
	}
}

func main() {
	t1 := &objeto{
		numero: 1,
		nombre: "Arroz",
		precio: 1000,
	}
	t2 := &objeto{
		numero: 2,
		nombre: "Azucar",
		precio: 1500,
	}
	t3 := &objeto{
		numero: 3,
		nombre: "Aguacate",
		precio: 2000,
	}
	t4 := &objeto{
		numero: 4,
		nombre: "Sal",
		precio: 800,
	}
	t5 := &objeto{
		numero: 5,
		nombre: "Carne",
		precio: 25000,
	}

	lista := &objetoLista{
		objetos: []*objeto{
			t1, t2, t3, t4, t5,
		},
	}

	// lista.agregarALista(t3)
	lista.imprimirLista()
	lista.agregarObjetos()

}
