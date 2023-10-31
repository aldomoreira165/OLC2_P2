# Interprete T-Swift

## Manual de Usuario

El uso del interprete es sencillo ya que cuenta con una interfaz bastante intuitiva. Al iniciar, encontrará una barra con los siguientes títulos:

* Archivo: este menú se compone de las siguientes opciones
    1. Crear archivo: Permite ingresar un nuevo archivo.
    2. Abrir archivo: Permite subir código desde el explorador de archivos de su computadora.
    3. Guardar archivo: Permite guardar el código en su computadora.
* Reportes: este menú permite lo siguiente
    1. Tabla de símbolos: Símbolos encontrados durante el análisis del código.
    2. Tabla de errores: En caso existan errores léxicos, sintácticos o semánticos durante la ejecución de su código estos podrán ser visualizados acá.
* Editar
* Acerca de

Interfaz de inicio.  
![Interfaz de inicio](https://imgur.com/v2zYP1B.png)

Subiendo un archivo: 
![Subiendo archivo](https://imgur.com/OwuYchn.png)

Ejecutando archivo
![Ejecutando archivo](https://imgur.com/2kl8q49.png)

Tabla de símbolos generada
![Tabla de símbolos](https://imgur.com/t2tyG9h.png)

Tabla de errores. En este caso se declaró dos veces una misma función.
![Tabla de errores](https://imgur.com/eM4KaY9.png)

Básicamente este es el flujo de la aplicación.

## Manual Técnico
El proyecto se basa en el lenguaje Golang, así como de la herramienta ANTLR en su versión 4. Se utilizó el patrón de diseño Visitor. Y para el frontend, se utilizó HTML5, CSS3 Y JS. Asimismo, se utilizó la librería Bootstrap y Graphviz.

Dado que ANTLR4 ofrece la facilidad de únicamente recorrer el árbol de análisis sintáctico el patron de diseño Visitor unicamente se encarga de visitar cada uno de los nodos para efectuar las operaciones necesarias para la correcta ejecución del código. 

Por lo cual, en función de los requerimientos el proyecto se moduló de la siguiente forma:

![Modulacion](https://imgur.com/j1xVRz0.png)

En donde cada uno de los archivos contiene el código necesario para ejecutar las operaciones dependiendo el tipo que sean. Finalmente, se tiene un modulo dedicado a la generación de código DOT para generar las imágenes en formato PNG de los distintos reportes mencionados anteriormente. 

![Estructuras](https://imgur.com/sN7oHKA.png)