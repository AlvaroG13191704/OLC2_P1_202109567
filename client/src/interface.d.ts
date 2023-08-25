

export interface Errors {
  line: number;
  column: number;
  type: string;
  message: string;
}

export interface SymbolTable {
  Id :          string // the name of the variable
	TypeSymbol:   string // the type of the symbol variable or function, vector, reference, struct
	TypeVariable: string // the type of the variable -> var or let, struct
	TypeData:     string // the type of the data -> Int, Float, String, Boolean, Character
	Value:        string[]
	ListParams:   string[]
	Mutating:     bool
	Line:         int
	Column:       int
}


export interface Analyzer {
  editorText: string;
  outputAnalysis: string;
  TableErrors: Errors[]
  TableSymbols: SymbolTable[]
}

export interface Response {
  result: string;
  errors: Errors[];
  symbols: SymbolTable[];
}