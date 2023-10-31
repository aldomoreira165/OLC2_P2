package reportes

import (
	"fmt"
	"interprete/lenguaje"
	"io/ioutil"
	"os"
	"os/exec"
)

func GenerarTabla(tabla *lenguaje.TablaSimbolos) {
	dotCode := generarTablaCode(tabla)
	dotFilePath := "tablaSimbolos.dot"

	err := ioutil.WriteFile(dotFilePath, []byte(dotCode), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo DOT:", err)
		return
	}
	defer os.Remove(dotFilePath)

	// Genera el archivo PNG
    pngFilePath := "./reportesPNG/tabla.png"
    cmd := exec.Command("dot", "-Tpng", dotFilePath, "-o", pngFilePath)
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error al generar el archivo PNG:", err)
        return
    }

    fmt.Println("Archivo PNG generado con éxito:", pngFilePath)
}

func generarTablaCode(tabla *lenguaje.TablaSimbolos) string {

	dotCode := "digraph G {\n"
	dotCode += "node[shape=record];\n"
	dotCode += "tbl [label=<<TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"
	dotCode += "<TR><TD>ID</TD><TD>Tipo de Símbolo</TD><TD>Tipo de Dato</TD><TD>Linea</TD><TD>Columna</TD></TR>\n"

	for _, simbolo := range tabla.Tabla {
		id := simbolo.ID
		tipoSimbolo := simbolo.TipoSimbolo
		tipoDato := simbolo.TipoDato
		linea := simbolo.Linea
		columna := simbolo.Columna

		dotCode += fmt.Sprintf("<TR><TD>%s</TD><TD>%s</TD><TD>%s</TD><TD>%d</TD><TD>%d</TD></TR>\n", id, tipoSimbolo, tipoDato, linea, columna)
	}
	dotCode += "</TABLE>>];\n"
	dotCode += "}"
	return dotCode
}