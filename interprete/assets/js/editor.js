const botonAbrirArchivo = document.getElementById("btn-abrir-archivo");
const botonCrearArchivo = document.getElementById("btn-crear-archivo");
const botonGuardarArchivo = document.getElementById("btn-guardar-archivo");
const btnEjecutar = document.getElementById("btn-ejecutar");
const btnTablaSimbolos = document.getElementById("btn-simbolos");
const btnTablaErrores = document.getElementById("btn-errores");

/*funcionalidad de editor de codigo*/

var Fuente = CodeMirror.fromTextArea(document.getElementById("fuente"), {
  mode: "text/x-java",
  theme: "dracula",
  lineNumbers: true,
  autoCloseTags: true,
});
Fuente.setSize("600", "579");

var Salida = CodeMirror.fromTextArea(document.getElementById("salida"), {
  mode: "text/x-java",
  theme: "dracula",
  lineNumbers: true,
  autoCloseTags: true,
});
Salida.setSize("600", "579");

/*acciones de los botones de la barra de menus*/

botonCrearArchivo.addEventListener("click", function () {
  Fuente.setValue("");
});

botonAbrirArchivo.addEventListener("click", function () {
  const input = document.createElement("input");
  input.type = "file";
  input.accept = ".tw";
  input.onchange = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.readAsText(file, "UTF-8");
    reader.onload = (readerEvent) => {
      const content = readerEvent.target.result;
      Fuente.setValue(content);
    };
  };
  input.click();
});

botonGuardarArchivo.addEventListener("click", function () {
  const fileName = window.prompt("Nombre del archivo:");
  if (fileName) {
    const textToSave = Fuente.getValue();
    const textToSaveAsBlob = new Blob([textToSave], { type: "text/plain" });
    const textToSaveAsURL = window.URL.createObjectURL(textToSaveAsBlob);
    const downloadLink = document.createElement("a");
    downloadLink.download = fileName + ".tw";
    downloadLink.innerHTML = "Descargar archivo";
    downloadLink.href = textToSaveAsURL;
    downloadLink.onclick = destroyClickedElement;
    downloadLink.style.display = "none";
    document.body.appendChild(downloadLink);
    downloadLink.click();
  }
});

function destroyClickedElement(event) {
  document.body.removeChild(event.target);
}

//funcionalidad para ejecutar archivo
btnEjecutar.addEventListener("click", function () {
  const codigo = Fuente.getValue();

  fetch("/enviar-codigo", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ contenido: codigo }),
  })
  .then((response) => {
    if (!response.ok) {
      return response.json(); // Leer el cuerpo de la respuesta como JSON
    }
    return response.json();
  })
  .then((data) => {
    if (data.error) {
      // Mostrar mensaje de error en la consola
      console.error("Error en el servidor:", data.error);

      // Mostrar mensaje de error en el textarea "Salida"
      Salida.setValue(data.error);
    } else {
      // Mostrar resultado en la interfaz
      Salida.setValue(data.salida);
    }
  })
  .catch((error) => {
    console.error("Error:", error);
  });
});