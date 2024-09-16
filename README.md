1. Crear mi archivo de modulos

> go mod init mimodulo

Cuando trabajamos con modulos si quiero que mis otros archivos los reconozca go debo incluirlos asi

> go run main.go operaciones.go utilidades.go

o hacer:

```sh
go build
./mimodulo.exe
```