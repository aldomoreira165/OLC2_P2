package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"interprete/Parser"
	"interprete/lenguaje"
	"interprete/reportes"
	"io/ioutil"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
)


type Respuesta struct {
	Salida  string `json:"salida"`
	Imagen  string `json:"imagen"`
	ImagenE string `json:"imagenE"`
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

	//generando tabla de simbolos
	tablaResult := visitor.(*lenguaje.Visitor).GetSymbolTable()
	reportes.GenerarTabla(tablaResult)

	//generando tabla de errores
	errorResult := visitor.(*lenguaje.Visitor).GetErrorTable()
	reportes.GenerarTablaErrores(errorResult)


	pngData, err := ioutil.ReadFile("./reportesPNG/tabla.png")
	if err != nil {
		fmt.Println("Error al leer el archivo PNG:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	pngDataE, err := ioutil.ReadFile("./reportesPNG/tablaErrores.png")
	if err != nil {
		fmt.Println("Error al leer el archivo PNG:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	//preparando respuesta
	respuesta := Respuesta{
		Salida:  out.(string),
		Imagen:  base64.StdEncoding.EncodeToString(pngData),
		ImagenE: base64.StdEncoding.EncodeToString(pngDataE),
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
