/*
● @autores: Jefferson Navarrete, Daniel Azcona
● @fecha: 24/11/2024
● @descripcion: Autonomo 2
● */
● package main
●
● import (
● "bufio"
● "fmt"
● "os"
● "strconv"
● "strings"
● )
●
● // Estructura de Cerveza que representa el inventario de cervezas
● type Cerveza struct {
● Nombre string
● Precio float64
● Cantidad int
● }
●
● // Estructura de Usuario para el registro y gestión de los
usuarios
● type Usuario struct {
● ID string
● Nombre string
● Email string
● Contraseña string
● Ciudad string
● }
●
● // Estructura de Pedido que representa un pedido realizado por un
usuario
● type Pedido struct {
● UsuarioID string
● Cerveza string
● Cantidad int
● Total float64
● Ciudad string
● }
●
● // Base de datos simulada en memoria: lista de cervezas y mapa de
usuarios
● var cervezas = make(map[string]*Cerveza)
● var usuarios = make(map[string]Usuario)
●
● // Función para agregar cervezas al inventario
● func agregarCerveza(nombre string, precio float64, cantidad int)
{
● cervezas[nombre] = &Cerveza{Nombre: nombre, Precio: precio,
Cantidad: cantidad}
● }
●
● // Función para mostrar todas las cervezas en el inventario
● func mostrarCervezas() {
● fmt.Println("Inventario de cervezas:")
● for _, cerveza := range cervezas {
● fmt.Printf("Nombre: %s, Precio: %.2f, Cantidad
disponible: %d\n",
● cerveza.Nombre, cerveza.Precio, cerveza.Cantidad)
● }
● }
●
● // Función para registrar un nuevo usuario
● func registrarUsuario(usuario Usuario) {
● usuarios[usuario.ID] = usuario
● }
●
● // Función para realizar un pedido por parte de un usuario
● func realizarPedido(usuarioID, nombreCerveza string, cantidad
int, ciudad string) (Pedido, error) {
● // Verificar si el usuario existe
● if _, usuarioExiste := usuarios[usuarioID]; !usuarioExiste {
● return Pedido{}, fmt.Errorf("Usuario con ID %s no
encontrado", usuarioID)
● }
●
● // Verificar si la cerveza existe y tiene stock suficiente
● cerveza, cervezaExiste := cervezas[nombreCerveza]
● if !cervezaExiste {
● return Pedido{}, fmt.Errorf("Cerveza con nombre %s no
encontrada", nombreCerveza)
● }
●
● // Verificar si hay suficiente stock
● if cerveza.Cantidad < cantidad {
● return Pedido{}, fmt.Errorf("Stock insuficiente para la
cerveza %s", cerveza.Nombre)
● }
●
● // Realizar el pedido
● total := cerveza.Precio * float64(cantidad)
● // Reducir el stock de la cerveza
● cerveza.Cantidad -= cantidad
●
● // Crear el pedido
● pedido := Pedido{
● UsuarioID: usuarioID,
● Cerveza: cerveza.Nombre,
● Cantidad: cantidad,
● Total: total,
● Ciudad: ciudad,
● }
●
● return pedido, nil
● }
●
● // Función para manejar la interacción con el cliente
● func realizarPedidoInteractivo() {
● reader := bufio.NewReader(os.Stdin)
●
● // Pedir información del usuario
● fmt.Print("Ingrese su nombre: ")
● nombre, _ := reader.ReadString('\n')
● nombre = strings.TrimSpace(nombre)
●
● fmt.Print("Ingrese su ciudad: ")
● ciudad, _ := reader.ReadString('\n')
● ciudad = strings.TrimSpace(ciudad)
●
● // Registrar al usuario en el sistema
● usuarioID := strconv.Itoa(len(usuarios) + 1) // Generar ID
único
● registrarUsuario(Usuario{ID: usuarioID, Nombre: nombre,
Ciudad: ciudad})
●
● fmt.Println("\nBienvenido,", nombre, "de", ciudad)
● fmt.Println("\nEste es nuestro inventario actual:")
● mostrarCervezas()
●
● // Seleccionar cerveza
● fmt.Print("\nIngrese el nombre de la cerveza que desea pedir:
")
● nombreCerveza, _ := reader.ReadString('\n')
● nombreCerveza = strings.TrimSpace(nombreCerveza)
●
● // Ingresar cantidad
● fmt.Print("Ingrese la cantidad que desea pedir: ")
● cantidadStr, _ := reader.ReadString('\n')
● cantidad, _ := strconv.Atoi(strings.TrimSpace(cantidadStr))
●
● // Procesar el pedido
● pedido, err := realizarPedido(usuarioID, nombreCerveza,
cantidad, ciudad)
● if err != nil {
● fmt.Println("Error al realizar el pedido:", err)
● } else {
● fmt.Printf("\nPedido realizado exitosamente:\nUsuario:
%s\nCerveza: %s\nCantidad: %d\nTotal: %.2f\nCiudad: %s\n",
● usuarios[pedido.UsuarioID].Nombre, pedido.Cerveza,
pedido.Cantidad, pedido.Total, pedido.Ciudad)
● }
●
● // Mostrar el inventario actualizado
● fmt.Println("\nInventario actualizado:")
● mostrarCervezas()
● }
●
● func main() {
● // Agregar cervezas al inventario
● agregarCerveza("IPA", 5.5, 100)
● agregarCerveza("Stout", 6.0, 50)
● agregarCerveza("Lager", 4.5, 200)
●
● // Ejecutar la interacción
● realizarPedidoInteractivo()
● }