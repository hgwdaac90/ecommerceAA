/*
@autores: Daniel Azcona, Jefferson Navarrete
@fecha: 09/12/2024
@descripcion: Etapa 3 - Implementación de Encapsulación, Manejo de Errores e Interfaces.
*/
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Interfaz para operaciones comunes
type Operaciones interface {
	Agregar()
	Mostrar()
}

// Usuario encapsulado
type Usuario struct {
	ci               string
	nombresCompletos string
}

// Getters y setters para Usuario
func (u *Usuario) GetCI() string {
	return u.ci
}

func (u *Usuario) SetCI(ci string) error {
	if ci == "" {
		return errors.New("El CI no puede estar vacío")
	}
	u.ci = ci
	return nil
}

func (u *Usuario) GetNombresCompletos() string {
	return u.nombresCompletos
}

func (u *Usuario) SetNombresCompletos(nombres string) error {
	if nombres == "" {
		return errors.New("Los nombres completos no pueden estar vacíos")
	}
	u.nombresCompletos = nombres
	return nil
}

// Producto encapsulado
type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func (p *Producto) GetNombre() string {
	return p.nombre
}

func (p *Producto) SetNombre(nombre string) error {
	if nombre == "" {
		return errors.New("El nombre del producto no puede estar vacío")
	}
	p.nombre = nombre
	return nil
}

func (p *Producto) GetPrecio() float64 {
	return p.precio
}

func (p *Producto) SetPrecio(precio float64) error {
	if precio <= 0 {
		return errors.New("El precio debe ser mayor a cero")
	}
	p.precio = precio
	return nil
}

// Orden encapsulada
type Orden struct {
	codigo    string
	usuario   Usuario
	productos []Producto
	total     float64
}

// Generación de código aleatorio para la orden
func GenerarCodigoOrden() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("ORD-%d", rand.Intn(100000))
}

// Agregar Usuario
func AgregarUsuario(usuarios *[]Usuario) {
	var ci, nombres string

	fmt.Println("Ingrese el CI del usuario:")
	fmt.Scanln(&ci)
	fmt.Println("Ingrese los nombres completos del usuario:")
	fmt.Scanln(&nombres)

	u := Usuario{}
	if err := u.SetCI(ci); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err := u.SetNombresCompletos(nombres); err != nil {
		fmt.Println("Error:", err)
		return
	}

	*usuarios = append(*usuarios, u)
	fmt.Println("Usuario agregado con éxito.")
}

// Realizar Pedido
func RealizarPedido(usuarios []Usuario, productos []Producto, ordenes *[]Orden) {
	var ci string
	fmt.Println("Ingrese el CI del usuario:")
	fmt.Scanln(&ci)

	// Buscar usuario
	var usuarioEncontrado *Usuario
	for i := range usuarios {
		if usuarios[i].GetCI() == ci {
			usuarioEncontrado = &usuarios[i]
			break
		}
	}
	if usuarioEncontrado == nil {
		fmt.Println("Usuario no encontrado.")
		return
	}

	// Crear orden
	codigo := GenerarCodigoOrden()
	orden := Orden{
		codigo:    codigo,
		usuario:   *usuarioEncontrado,
		productos: productos, // Simulamos una lista de productos
		total:     CalcularTotal(productos),
	}
	*ordenes = append(*ordenes, orden)

	fmt.Printf("Pedido generado con éxito. Código de orden: %s\n", codigo)
	fmt.Println("Para realizar su pago, realice una transferencia bancaria a la cuenta 12345 - Banco Pichincha.")
	fmt.Println("Envíe el comprobante y el número de orden a los números +593997275431 o +593967077066 (Quito - Ecuador).")
}

// Mostrar Pedido
func MostrarPedido(ordenes []Orden) {
	var codigo string
	fmt.Println("Ingrese el código de la orden:")
	fmt.Scanln(&codigo)

	// Buscar orden
	for _, orden := range ordenes {
		if orden.codigo == codigo {
			fmt.Printf("Detalles de la orden:\nCódigo: %s\nUsuario: %s\nTotal: $%.2f\n", orden.codigo, orden.usuario.GetNombresCompletos(), orden.total)
			return
		}
	}
	fmt.Println("Orden no encontrada.")
}

// Calcular Total de una lista de productos
func CalcularTotal(productos []Producto) float64 {
	total := 0.0
	for _, p := range productos {
		total += p.GetPrecio()
	}
	return total
}

func main() {
	usuarios := []Usuario{}
	productos := []Producto{
		{"Cerveza IPA", 5.50, 10},
		{"Cerveza Stout", 6.00, 8},
	}
	ordenes := []Orden{}

	for {
		var opcion int
		fmt.Println("1. Agregar Usuario")
		fmt.Println("2. Realizar Pedido")
		fmt.Println("3. Mostrar Pedido")
		fmt.Println("4. Salir")
		fmt.Scanln(&opcion)

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Se ha producido un error:", r)
			}
		}()

		switch opcion {
		case 1:
			AgregarUsuario(&usuarios)
		case 2:
			RealizarPedido(usuarios, productos, &ordenes)
		case 3:
			MostrarPedido(ordenes)
		case 4:
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
