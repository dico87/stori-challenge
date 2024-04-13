# Diego Cortes - Stori Challenge
## Summary
Este proyecto consiste en realizar un procesador de transacciones que permita al usuario generar una informacion de unas transacciones extraidas de un archivo en formato CSV

## Estructura del Proyecto
Este proyecto tiene una estructura orientada a [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)

Tenemos una aplicación con la siguiente estructura:

* **cmd**
  * **transactions** 
    * **main.go** - `Es el archivo de ejecucion de la aplicacion`
* **internal**
  * **repository** - `Este directorio contiene los repositorios de la aplicacion para en esta ocasion almacenar en base de datos`
    * **entities** - `Entidades que mapean la tablas de base de datos`
  * **transactions** - `Contiene toda la logica de la aplicacion`
    * **domain** - `Entidades de dominio de la aplicacionm no son entidades planas si no entidades enriquecidas que permiten encapsular logica de negocio en ellas representando todo el dominio en ellas`
    * **mocks** - `Contiene algunos mocks de interfaces para testing`
    * **reader** - `Logica sobre la lectura de fuentes de informacion, en nuestro caso archivos CSV`
    * **sender** - `Logica sobre envio de la informacion en nuestro caso Correo Electronico`
    * **transactions.go** - `Archivo que orquesta el procesamiento de informacion usando domain, sender y reader`
    
## Testing
Cada archivo que contiene logica y proceso en la aplicacion contiene sus test respectivos.

Se realiza el uso de [Test Suite](https://pkg.go.dev/github.com/stretchr/testify/suite) el cual es un componente de la libreria `testify` que permite un uso mas ordenado de los tests.

### Covertura
En este momento la aplicacion cuenta con 89.9% de covertura de codigo, este porcentaje no es un 100% debido a aquellos casos con errores en base de datos y envio de correos a los cuales implementar un mock dependee de librerias externas y no generan un valor a la hora del resultado final.

## Anotaciones
1. Se realiza uso de inyeccion de dependencias haciando todas las declaraciones desde el archivo `main.go` donde see crean todas las dependencias necesarias y se van armando los componentes de la aplicación inyectando todas las dependencias necesitas que requieren. Esto permite hacer testing de una manera mas facil.
2. Se utilizan interfaces de comunicacion entre el procesador principal `transactions.go` y sus componentes como `sender` y `reader`. Esto permite que se pueda extender la funcionalidad creando nuevas implementaciones tanto para lectura como envio sin modificar la logica de la aplicación.
3. Se utiliza el ORM gorm V2 para toda la gestion de Base de Datos
4. El motor utilizado para base de datos se SQLite manejandolo en un archivo `transactions.db` que sera creado al momento de comenzar a generar transacciones.

## Variables de Entorno
Como la aplicacion maneja datos sensibles como el usuario y contraseña del correo electronico, estas configuraciones se realizan por medio de variable de entorno las cuales son: 

1. `email_password`
2. `email_smtp_server`
3. `email_subject`
4. `email_to`
5. `email_user`

Estas deben ser establecidas antes de ejecutar la aplicación

## Como ejecutar la aplicacion

### Docker
Se incluye el archivo `Dockerfile` con el cual se puede generar una imagen docker y ejecutarla, se debe realizar la modificacion de las variables de entorno mencionadas anteriormente en el archivo `Dockerfile` para que la aplicacion envie correctamente los correos electronicos