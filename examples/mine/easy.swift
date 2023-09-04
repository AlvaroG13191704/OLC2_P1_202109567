
func fibonacci(_ n: Int) -> Int {
    if n > 1 {
        return fibonacci(n - 1) + fibonacci(n - 2)
    } else if n == 1 {
        return 1
    } else if n == 0 {
        return 0
    } else {
        print("error")
        return 0
    }
}

print("Debería ser 55")
print(fibonacci(10)) 

func ackerman(_ m: Int, _ n: Int) -> Int {
    if m == 0 {
        return n + 1
    } else if m > 0 && n == 0 {
        return ackerman(m - 1, 1)
    } else {
        return ackerman(m - 1, ackerman(m, n - 1))
    }
}

print("Debería ser 125")
print(ackerman(3,4))

func Hanoi(_ discos: Int, _ origen: Int, _ auxiliar: Int, _ destino: Int) {
    if discos == 1 {
        print("Mover disco de", origen, "a", destino)
    } else {
        Hanoi(discos - 1, origen, destino, auxiliar)
        print("Mover disco de", origen, "a", destino)
        Hanoi(discos - 1, auxiliar, origen, destino)
    }
}

print("Hanoi")
Hanoi(3, 1, 2, 3)


var vector: [Int] = [1, 2, 3, 4, 5]

// duplicar valores de un vector
func duplicar(_ arr: inout [Int]) {
    var i = 0
    while i < arr.count {
        arr[i] += arr[i]
        i += 1
    }
}

print("==============ANTES=================")
for i in 0...vector.count-1 {
    print(vector[i])
}

print("==============Después=================")
print("Debería ser [2, 4, 6, 8, 10]")

duplicar(&vector)

for i in 0...vector.count-1 {
    print(vector[i])
}



// Función para intercambiar dos elementos en un arreglo
func intercambiar(_ a: inout [Int], _ i: Int, _ j: Int) {
    let aux = a[i]
    a[i] = a[j]
    a[j] = aux
}

// Algoritmo de ordenamiento por selección (Selection Sort)
func ordSeleccion(_ arr: inout [Int]) {
    var i = 0
    var j = 0
    var indiceMenor = 0
    let n = arr.count
    
    while i < (n - 1) {
        indiceMenor = i
        j = i + 1
        while j < n {
            if arr[j] < arr[indiceMenor] {
                indiceMenor = j
            }
            j += 1
        }
        
        if i != indiceMenor {
            intercambiar(&arr, i, indiceMenor)
        }
        i += 1
    }
}
var arr2: [Int] = [40, 21, 1, 3, 14, 4]
ordSeleccion(&arr2)

for i in 0...arr2.count-1{
    print(i)
}




// Función para imprimir un arreglo
func printArray(_ msg: String, _ arr: [Int]) {
    var out = " ["
    for i in 0...arr.count-1 {
        if i == arr.count - 1 {
            out += String(arr[i])
        } else {
            out += String(arr[i]) + ", "
        }
    }
    out += "] "
    print(msg + out)
}

var arr1: [Int] = [8, 4, 6, 2]
printArray("entrada: ", arr1)



func printWord(_ word: String) {
    print(word)
}

printWord("Hola mundo")