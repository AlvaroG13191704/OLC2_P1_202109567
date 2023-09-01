# Manual Técnico 
## Introducción
El presente documento tiene como objetivo describir la estructura del proyecto, así como las tecnologías utilizadas para su desarrollo. La finalidad de este proyecto fue construir un interprete del lenguaje de programación Swift, donde se le nombro T-swift. El interprete fue desarrollado en el lenguaje de programación Go, utilizando la herramienta ANTLR para la generación del lexer y parser del lenguaje.
## Estructura del proyecto
El proyecto se encuentra dividido en dos carpetas principales: `frontend` y `backend`. La primera contiene el código del cliente, mientras que la segunda contiene el código del servidor. A continuación se describen las carpetas y archivos más importantes de cada una de ellas.
### Frontend
- **src**: Contiene todo el código fuente del cliente. Asi como los componentes, estilos, imágenes, etc.
- **public**: Contiene el archivo `index.html` que es el punto de entrada de la aplicación.
- **package.json**: Contiene la información del proyecto, así como las dependencias necesarias para su ejecución.

### Backend
- **interpreter**: Contiene el código del intérprete de la aplicación.
- **routes**: Contiene los archivos que definen las rutas del servidor.
- **server.go**: Contiene el código del servidor.
- **go.mod**: Contiene la información del proyecto, así como las dependencias necesarias para su ejecución.
- **go.sum**: Contiene la información de las dependencias del proyecto.
- **SFGrammar.g4**: Contiene la gramática del lenguaje.
- **SFLexer.g4**: Contiene el lexer del lenguaje.

## Tecnologías utilizadas
### Frontend - Cliente
- **Typescript**: Lenguaje de programación. Este lenguaje fue utilizado para el desarrollo del cliente de la aplicación.
- **Sveltekit**: Framework de desarrollo web. Este framework es un framework de desarrollo web basado en componentes, el cual utiliza el lenguaje de programación Typescript. Este framework fue utilizado para el desarrollo del cliente de la aplicación. El propio framework se encarga del manejo de los componentes y la comunicación con el servidor.
- **TailwindCSS**: Framework de estilos. Este framework fue utilizado para el desarrollo de los estilos de la aplicación.

### Backend
- **Go**: Lenguaje de programación. Este lenguaje fue utilizado para el desarrollo del servidor de la aplicación.
- **Fiber**: Framework web. Este framework fue utilizado para el desarrollo del servidor de la aplicación. El propio framework se encarga del manejo de las rutas y la comunicación con el cliente.
- **ANTLR4**: Herramienta para la generación de lexer y parser. Esta herramienta fue utilizada para la generación del lexer y parser del lenguaje. El lexer y parser fueron generados a partir de la gramática del lenguaje. Esta herramienta utiliza dos patrone de diseño: Visitor y Listener.

## Endpoint del servidor
El servidor cuenta con un único endpoint, el cual recibe el código fuente del programa a interpretar y retorna la tabla de símbolos, la tabla de errores y los outputs del programa. El endpoint se encuentra en la ruta `/analyze`. El endpoint recibe text/plain:
```go
curl --location 'http://127.0.0.1:3000/analyze' \
--header 'Content-Type: text/plain' \
--data 'struct Persona {
    var Nombre: String
    var edad = 0

    func saludar() {
        print("hola mi nombre es: ", self.Nombre)
    }
}

var alvaroP = Persona(Nombre:"Alvaro", edad: 20 )

alvaroP.saludar()'
```
La respuesta  es un json con la tabla de símbolos, la tabla de errores y los outputs del programa:
```json
{
    "dot":"svg data",
    "errors": ["LISTA DE ERRORES"],
    "result": "RESULTADO DEL PROGRAMA",
    "symbol": [
        {
            "Id": "alvaroP",
            "TypeSymbol": "variable",
            "TypeVariable": "var",
            "TypeData": "Struct",
            "StructOf": "Persona",
            "Value": {
                "Nombre": {
                    "Id": "Nombre",
                    "TypeSymbol": "variable",
                    "TypeVariable": "var",
                    "TypeData": "String",
                    "StructOf": "",
                    "Value": {
                        "Value": "Alvaro"
                    },
                    "ListParams": null,
                    "Mutating": false,
                    "Line": 2,
                    "Column": 4
                },
                "edad": {
                    "Id": "edad",
                    "TypeSymbol": "variable",
                    "TypeVariable": "var",
                    "TypeData": "Int",
                    "StructOf": "",
                    "Value": {
                        "Value": 20
                    },
                    "ListParams": null,
                    "Mutating": false,
                    "Line": 3,
                    "Column": 4
                },
                "saludar": {
                    "Id": "saludar",
                    "TypeSymbol": "function",
                    "TypeVariable": "void",
                    "TypeData": "function",
                    "StructOf": "",
                    "Value": {
                        "RuleIndex": 1
                    },
                    "ListParams": null,
                    "Mutating": false,
                    "Line": 5,
                    "Column": 4
                }
            },
            "ListParams": null,
            "Mutating": false,
            "Line": 10,
            "Column": 0
        }
    ]
}
```

## Patron de diseño utilizado - Visitor
El patrón de diseño utilizado para la generación del lexer y parser fue el patrón de diseño Visitor. Este patrón de diseño permite separar el algoritmo de un objeto de su representación. De esta forma, se puede agregar nuevas operaciones a la estructura de objetos sin modificar dicha estructura. En este caso, se utilizo este patrón de diseño para agregar las acciones que se deben realizar al momento de visitar un nodo del árbol de sintaxis abstracta. De esta forma, se puede agregar nuevas acciones sin modificar la estructura del árbol de sintaxis abstracta.

## Estructuras importanes para el manejo de toda la lógica del interprete

### Struct para manejar los símbolos de la tabla de símbolos
Este struct se utiliza para manejar los símbolos de la tabla de símbolos. Este struct contiene la información del símbolo, como el id, el tipo de símbolo, el tipo de variable, el tipo de dato, el struct al que pertenece, el valor, los parámetros, si es mutante, la línea y la columna donde se encuentra. Funciona para toda variable, vector, función, struct, matriz, etc.
```go
type SymbolTable struct {
	Id           string
	TypeSymbol   string
	TypeVariable string
	TypeData     string 
	StructOf     string 
	Value        interface{}
	ListParams   interface{}
	Mutating     bool
	Line         int
	Column       int
}
```

### Struct para manejar los errores
Este struct se utiliza para manejar los errores sintácticos y semánticos. Este struct contiene la información del error, como la línea, la columna, el mensaje y el tipo de error.
```go
type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}
```

### Struct para manejar las sentencias de transferencia
Este struct se utiliza para manejar las sentencias de transferencia. Este struct contiene la información de la sentencia de transferencia, como el tipo de sentencia, la línea y la columna donde se encuentra. Sirve para el manejo del "continue" y "break".
```go
type LoopContext struct {
	TypeLoop      string
	ContinueFound bool
	BreakFound    bool
}
```
### Struct para manejar el self
Este struct se utiliza para manejar el self. Este struct contiene la información del self, como la instancia de cada struct, el tipo de struct que pertenece la variable, nombre de la función activa y una variable booleana para saber si se esta llamando a una función o no.
```go
type SelfStruct struct {
	VarId        string
	StructOf     string
	FunctionName string
	FunctionCall bool
}
```

### Struct del Visitor, donde nacen todos los métodos para recorrer el AST
- **SymbolStack**: Es una pila de mapas de símbolos, donde se almacenan los símbolos de cada función, struct, etc. Funciona como el ambito de todo el programa, cuentan con metodos para agregar, buscar, eliminar y actualizar símbolos y crear o eliminar nuevos ambitos.
- **TableSymbol**: Es una lista de símbolos, donde se almacenan los símbolos declarados en el programa. Su función principal es para la generación de la tabla de símbolos.
- **Outputs**: Es una lista de strings, donde se almacenan los outputs del programa o sea donde un print fue llamado.
- **Errors**: Es una lista de errores, donde se almacenan los errores sintácticos y semánticos del programa. Su función principal es para la generación de la tabla de errores.
- **loopContexts**: Es una lista de LoopContext, donde se almacenan los contextos de los loops, para el manejo de los "continue" y "break".
- **ReturnValue**: Es una variable interface{}, donde se almacena el valor de retorno de una función.
- **IsReturn**: Es una variable booleana, donde se almacena si una función tiene un return o no.
- **FirstPass**: Es una variable booleana, donde se almacena si es el primer recorrido del AST o no.
- **FunctionContext**: Es una lista de strings, donde se almacenan los nombres de las funciones que se estan ejecutando.
- **SelfStructs**: Es un mapa de SelfStruct, donde se almacenan los selfs de cada struct. Su función principal es para el manejo del self.
```go
type Visitor struct {
	parser.BaseSFGrammarVisitor
	SymbolStack []map[string]SymbolTable
	TableSymbol []SymbolTable
	Outputs     []string
	Errors      []Error
	loopContexts    []LoopContext
	ReturnValue     interface{}
	IsReturn        bool
	FirstPass       bool
	FunctionContext []string
	SelfStructs map[string]SelfStruct
}
```
