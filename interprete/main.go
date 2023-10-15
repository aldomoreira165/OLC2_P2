package main

import (
	"encoding/json"
	"fmt"
	"interprete/Parser"
	"interprete/lenguaje"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
)


type Respuesta struct {
	Salida  string `json:"salida"`
}

type CodigoEnviado struct {
	Contenido string `json:"contenido"`
}


func manejarEnviarcodigo(w http.ResponseWriter, r *http.Request) {
	// Decodificar el cuerpo JSON en una estructura CodigoEnviado
	var codigo CodigoEnviado
	err := json.NewDecoder(r.Body).Decode(&codigo)
	if err != nil {
		http.Error(w, "Error al leer el contenido JSON", http.StatusBadRequest)
		return
	}
	code := codigo.Contenido
	input := antlr.NewInputStream(code)
	lexer := parser.NewSwiftLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSwiftGrammarParser(tokens)
	p.BuildParseTrees = true
	tree := p.S()
	visitor := lenguaje.NewVisitor()
	out := visitor.Visit(tree)
	fmt.Println("Resultado: ", out)

	//preparando respuesta
	respuesta := Respuesta{
		Salida:  out.(string),
	}

	// enviando respuesta al cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func main() {
	fs := http.FileServer(http.Dir("."))
	http.HandleFunc("/enviar-codigo", manejarEnviarcodigo)
	http.Handle("/", fs)
	http.ListenAndServe("localhost:3000", nil)
}

//antlr4 -Dlanguage=Go -o parser -package parser -visitor *.g4
