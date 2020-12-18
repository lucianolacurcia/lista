// Package lista provides Librería básica del TAD lista
package lista

import "fmt"

type Nodo_lista struct {
	siguiente *Nodo_lista
	anterior  *Nodo_lista
	elemento  uint
}

type Rep_lista struct {
	inicio *Nodo_lista
	ultimo *Nodo_lista
	actual *Nodo_lista
}

type Lista *Rep_lista

func Crear_nodo(elemento uint, anterior *Nodo_lista, siguiente *Nodo_lista) *Nodo_lista {
	nuevo := new(Nodo_lista)
	nuevo.elemento = elemento
	nuevo.siguiente = siguiente
	nuevo.anterior = anterior
	if siguiente != nil {
		siguiente.anterior = nuevo
	}
	if anterior != nil {
		anterior.siguiente = nuevo
	}
	return nuevo
}

func Crear_lista() Lista {
	res := new(Rep_lista)
	res.inicio = nil
	res.ultimo = nil
	res.actual = nil
	return res
}

func Copiar_lista(lst Lista) Lista {
	res := Crear_lista()
	var cursor_lst *Nodo_lista = lst.inicio
	var ptr_nuevo **Nodo_lista = &res.inicio
	var anterior *Nodo_lista = nil
	for cursor_lst != nil {
		var nuevo *Nodo_lista = Crear_nodo(cursor_lst.elemento, anterior, nil)
		*ptr_nuevo = nuevo
		anterior = nuevo
		res.ultimo = nuevo
		ptr_nuevo = &(nuevo.siguiente)
		cursor_lst = cursor_lst.siguiente
	}
	return res
}

func Insertar_al_inicio(elem uint, lst Lista) {
	nuevo := Crear_nodo(elem, nil, lst.inicio)
	lst.inicio = nuevo
	if lst.ultimo == nil {
		lst.ultimo = nuevo
	}
	lst.actual = nil
}

func Insertar_al_final(elem uint, lst Lista) {
	nuevo := Crear_nodo(elem, lst.ultimo, nil)
	if lst.inicio == nil {
		lst.inicio = nuevo
	} else {
		lst.ultimo.siguiente = nuevo
	}
	lst.ultimo = nuevo
	lst.actual = nil
}

func Remover(elem uint, lst Lista) {
	var cursor_lst **Nodo_lista = &lst.inicio
	var anterior *Nodo_lista = nil
	var a_borrar *Nodo_lista = nil
	for *cursor_lst != nil {
		if (*cursor_lst).elemento == elem {
			a_borrar = *cursor_lst
			*cursor_lst = a_borrar.siguiente
			if *cursor_lst != nil {
				(*cursor_lst).anterior = anterior
			} else {
				lst.ultimo = a_borrar.anterior
				if lst.ultimo == nil {
					lst.inicio = nil
				}
			}
			lst.actual = nil
			return
		}
		anterior = *cursor_lst
		cursor_lst = &((*cursor_lst).siguiente)
	}

}

func Remover_al_inicio(lst Lista) {
	if !Es_vacia_lista(lst) {
		if lst.inicio != nil {
			a_borrar := lst.inicio
			lst.inicio = a_borrar.siguiente
			if lst.inicio != nil {
				lst.inicio.anterior = nil
			} else {
				lst.ultimo = nil
			}
			lst.actual = nil
		}
	}
}

func Remover_al_final(lst Lista) {
	if !Es_vacia_lista(lst) {
		if lst.ultimo != nil {
			a_borrar := lst.ultimo
			lst.ultimo = a_borrar.anterior
			if lst.ultimo != nil {
				lst.ultimo.siguiente = nil
			} else {
				lst.inicio = nil
			}
			lst.actual = nil
		}
	}
}

func Es_vacia_lista(lst Lista) bool {
	return (lst.inicio == nil || lst.ultimo == nil)
}

func Pertenece_a_lista(elem uint, lst Lista) bool {
	cursor := lst.inicio
	for (cursor != nil) && (cursor.elemento != elem) {
		cursor = cursor.siguiente
	}
	return cursor != nil
}

func Destruir_lista(lst Lista) {
	fmt.Printf("Golang tiene garbage collector :) \n")
}

func Cantidad_elementos_lista(lst Lista) uint {
	var cantidad uint = 0
	cursor := lst.inicio
	for cursor != nil {
		cantidad++
		cursor = cursor.siguiente
	}
	return cantidad
}

func Comienzo_lista(lst Lista) {
	lst.actual = lst.ultimo
}

func Final_lista(lst Lista) {
	lst.actual = lst.ultimo
}

func Anterior_lista(lst Lista) uint {
	elem := lst.actual.elemento
	lst.actual = lst.actual.anterior
	return elem
}

func Siguiente_lista(lst Lista) uint {
	elem := lst.actual.elemento
	lst.actual = lst.actual.siguiente
	return elem
}

func Existe_actual_lista(lst Lista) bool {
	return lst.actual != nil
}

func Actual_lista(lst Lista) uint {
	return lst.actual.elemento
}

func Primero_lista(lst Lista) uint {
	return lst.inicio.elemento
}

func Ultimo_lista(lst Lista) uint {
	return lst.ultimo.elemento
}

func Imprimir_lista(lst Lista) {
	Comienzo_lista(lst)
	for Existe_actual_lista(lst) {
		fmt.Printf("%d ", Actual_lista(lst))
		Siguiente_lista(lst)
	}
}
