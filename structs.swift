// Matrix declaration

var mtx1: [[Int]] = [[1], [4, 5], [7, 8, 9]]

var mtx2 : [[[Int]]] = [[[1,2,3],[4,5,6],[7,8,9]], 
                        [[10,11,12],[13,14,15],[16,17,18]], 
                        [[19,20,21],[22,23,24],[25,26,27]]]

// Using repeating

var matrix2 : [[[Int]]] = [[[Int]]] (repeating: [[Int]] (repeating: [Int] (repeating: 0, count:2), count:3), count:4)

var matrix0 : [[[String]]] = [[[String]]] (repeating: [[String]] (repeating: [String] (repeating:"OLC2", count:2), count:1), count:3)


print(vec1[0], "======") //imprime 20
vec1[0] = mtx1[0][0] //cambia 20 por 1
print(vec1[0], "======") //imprime 1

print(mtx1[1][1], "^^^^") //imprime 5
mtx1[1][1] = 10 //cambia 5 por 10 
print(mtx1[1][1], "^^^^") //imprime 10

print(mtx1[0][0]) //imprime 1 

print(mtx2[0][1][2], "-----") //imprime 6
mtx2[0][1][2] = 10 //cambia 6 por 10
print(mtx2[0][1][2], "-----") //imprime 10

//error indices fuera de rango - error
mtx1[100][100] = 10

// Structs

struct Persona{
    var Nombre: String?
    var edad = 0
}

// struct con funciones 
struct Avion {
    var pasajeros = 0 
    var velocidad = 100
    var nombre: String?
    // var piloto: Persona

    // metodo dentro de struct 

    mutating func frenar(){
        print("Frenando")
        //al ser mutable sí afecta al struct 
        // self.velocidad = 0
    }
        
    // funcion inmutable 
    func mostrarVelocidad(){
        print("Velocidad", self.velocidad) 
    }
}

var avioneta = Avion( nombre: "78496", piloto: Persona(Nombre: "Joel", edad: 43 ))

// acceso a un atributo
print(avioneta.pasajeros)
// modificion de un atributo
avioneta.pasajeros = 5

print("Velocidad", avioneta.pasajeros) // llamada de la funcion avioneta.mostrarVelocidad()
// copia de structs por valor
var avioneta2 = avioneta
avioneta2.pasajeros = 0
//imprime: avioneta.pasajeros: 5 
print("avioneta.pasajeros:",avioneta.pasajeros) 
//avioneta2.pasajeros 0 
print("avioneta2.pasajeros:", avioneta2.pasajeros)
print("avioneta.piloto.Nombre:", avioneta2.piloto.Nombre )

struct Fruta{
    let nombre: String = "pera" 
    var precio: Int?
}
// solo se puede definir precio en el constructor 
// si se llega a definir nombre será un error 
var pera = Fruta(precio: 10)


struct Persona {
    var name: String var age: Int
}
var personas = [Persona]()
// se agregan valores al arra 
personas.append(Persona(Nombre: "Celeste", edad: 23)) 
personas.append(Persona(Nombre: "Roel", edad: 32)) 
personas.append(Persona(Nombre: "Flor", edad: 17))

var persona1 = personas[0] 
persona1.Nombre = "Nancy"

print(persona1.Nombre) //imprime Nancy 
print(personas[0].Nombre) //imprime Celeste 
// se modifica un array
personas[1].edad = 26

// otras formas permitidas 
struct Distro {
    var Nombre: String?
    var Version: String? 
}
var Distros = [
    Distro(Nombre: "Ubuntu", Version: "22.04"), 
    Distro(Nombre: "Fedora", Version: "38"), 
    Distro(Nombre: "OpenSUSE", Version: "Leap 15")
]

print(Distros[0].Nombre) // Imprime Ubuntu 
print(Distros[1].Version) // Imprime 13