package reportes

import (
	"fmt"
	"interprete/lenguaje"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func GenerarTablaErrores(tabla *lenguaje.TablaErrores) {
	dotCode := generarTablaCodeErrores(tabla)
	dotFilePath := "tablaErrores.dot"

	err := ioutil.WriteFile(dotFilePath, []byte(dotCode), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo DOT:", err)
		return
	}
	defer os.Remove(dotFilePath)

	// Genera el archivo PNG
    pngFilePath := "./reportesPNG/tablaErrores.png"
    cmd := exec.Command("dot", "-Tpng", dotFilePath, "-o", pngFilePath)
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error al generar el archivo PNG:", err)
        return
    }

    fmt.Println("Archivo PNG generado con éxito:", pngFilePath)
}

func generarTablaCodeErrores(tabla *lenguaje.TablaErrores) string {

	dotCode := "digraph G {\n"
	dotCode += "node[shape=record];\n"
	dotCode += "tbl [label=<<TABLE BORDER=\"1\" CELLBORDER=\"1\" CELLSPACING=\"0\">\n"
	dotCode += "<TR><TD>Descripción</TD><TD>Línea</TD><TD>Columna</TD></TR>\n"

	for _, error := range tabla.Tabla {
		descr := error.Description
		linea := error.Linea
		columna := error.Columna

		//convertir linea y columna a string
		lineaStr := strconv.Itoa(linea)
		columnaStr := strconv.Itoa(columna)

		dotCode += fmt.Sprintf("<TR><TD>%s</TD><TD>%s</TD><TD>%s</TD></TR>\n", descr, lineaStr, columnaStr)
	}
	dotCode += "</TABLE>>];\n"
	dotCode += "}"
	return dotCode
}