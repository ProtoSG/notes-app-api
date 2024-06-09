# API Aplicación de Notas

## Descripción

La Aplicación de Notas es una API desarrollada en Go, diseñada para gestionar notas de manera eficiente. Permite a los usuarios crear, editar y eliminar notas, así como etiquetarlas y organizarlas por categorías. Además, los usuarios pueden compartir sus notas con otros usuarios.

## Características

- Crear Notas: Los usuarios pueden crear nuevas notas con contenido personalizado.
- Crear Nota- Editar Notas: Los usuarios pueden editar notas existentes para actualizar su contenido.
- Eliminar Notas: Los usuarios pueden eliminar notas que ya no necesiten.
- Etiquetado: Los usuarios pueden etiquetar sus notas para una mejor organización y búsqueda.
- Categorías: Las notas se pueden organizar en categorías, permitiendo una estructura más ordenada.
- Compartir Notas: Los usuarios pueden compartir sus notas con otros usuarios, facilitando la colaboración.

## Requisitos

- Go 1.18 o superior
- MongoDB

## Instalación

1. Clonar el repositorio:

```bash
git clone https://github.com/ProtoSG/notes-app-api
```

2. Navegar al directorio del proyecto:

```bash
cd notes-app-api
```

3. Instalar las dependencias:

```bash
go mod tidy
```

## Configuración

Configurar las variables de entorno necesarias (por ejemplo, para la base de datos y JWT):

```bash
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=notasApp
JWT_SECRET_KEY=tu_secreto
```

## Uso

1. Ejecutar la aplicación:

```bash
go run cmd/main.go
```

2. La API estará disponible en `http://localhost:8080`.

## Endpoints

### Autenticación

- **POST /register**: Registrar un nuevo usuario.
- **POST /login**: Iniciar sesión y obtener un token JWT.

### Notas

- **GET /notes**: Obtener todas las notas del usuario autenticado.
- **POST /notes**: Crear una nueva nota.
- **GET /notes/{id}**: Obtener una nota específica.
- **PUT /notes/{id}**: Actualizar una nota específica.
- **DELETE /notes/{id}**: Eliminar una nota específica.

### Etiquetas y Categorías

- **POST /notes/{id}/tags**: Añadir etiquetas a una nota.
- **POST /notes/{id}/categories**: Añadir una nota a una categoría.

### Compartir Notas

- **POST /notes/{id}/share**: Compartir una nota con otro usuario.

## Estructura del Proyecto

```go
notes-app-api/
|-- cmd/
  ├── main.go
  |-- adapters/
  | ├── note.adapter.go
  | ├── user.adapter.go
  |
  ├── controllers/
  │ ├── auth.controller.go
  │ ├── note.controller.go
  │
  ├── models/
  │ ├── note.model.go
  │ ├── user.model.go
  │
  ├── routes/
  │ ├── auth.routes.go
  │ ├── notes.routes.go
  │
  ├── middleware/
  │ ├── auth.middleware.go
  │
  ├── utils/
  │ ├── jwt.go
  │ ├── responses.go
  | |-- server.go
  │
  |-- db/
  | ├── connection.go
  |
  └── config/
    └── config.go
```
