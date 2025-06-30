# ml-item-detail

Proyecto que implementa un detalle de producto inspirado en MercadoLibre.

El foco principal está en el backend, desarrollado en Go con SQLite como base de datos.  
El entorno está preparado para ser **plug & play** mediante `docker-compose`.

La mayor parte del frontend fue desarrollado con asistencia de v0, dado que el énfasis principal está en el backend.  
Se diseñó un diagrama de flujo de modelos para clarificar cómo se estructuran y se deben poblar los datos. (https://dbdiagram.io/d/6860af80f413ba350858eac6)

El backend expone dos endpoints principales:  
- `items/:id` — requerido para el challenge, que devuelve el detalle de un ítem específico.  
- `items/` — endpoint adicional que reutiliza el repositorio existente, permitiendo filtros como nombre y categoría para búsqueda y listado. Se encuentra fuera del scope del challenge pero se puede consultar utilizando curl.

Además, se incorporó un action para ejecutar el linter.
Se implementaron tests unitarios para cada repositorio, usecase, entrypoint. También se definieron las estructuras de datos y sus relaciones correspondientes.

---

## 🛠️ Requisitos

- Docker  
- Docker Compose

---

## 🚀 Levantar el proyecto

```bash
docker-compose down --volumes --remove-orphans # opcional
docker-compose build --no-cache && docker-compose up
```
Estos comandos:
- Opcionalmente -> Elimina volúmenes previos y remueve contenedores huérfanos para inicializar limpiamente 
- Hace un build sin usar cache, asegurando que se tomen todos los cambios recientes
- Levanta los dos proyectos

Cómo funciona
- Todos los puertos y variables de entorno están definidos en docker-compose.yml.
- El backend utiliza una base de datos SQLite persistida en ./data.
- Al levantar el contenedor, si SEED=true, se ejecuta el seeder:
    - Si es la primera ejecución, se insertan datos de prueba automáticamente.
    - En siguientes ejecuciones, se evita la duplicación de datos.
- Luego de levantados los servicios. Se puede acceder desde http://localhost:3000

📁 Estructura del proyecto
```bash
ml-item-detail/
|-- backend
    |-- config
    |   |-- config.go
    |   `-- dbconn.go
    |-- data
    |   `-- database.db
    |-- docs
    |-- packages
    |   `-- dbconn
    |       `-- sqlite.go
    |-- src
    |   |-- app
    |   |   |-- router
    |   |   |-- web
    |   |   `-- application.go
    |   |-- core
    |   |   |-- contracts
    |   |   |-- entities
    |   |   |-- providers
    |   |   |-- seed
    |   |   |-- usecases
    |   |   `-- utils
    |   |-- entrypoints
    |   |   |-- rest
    |   |   |-- items.go
    |   |   `-- ping.go
    |   |-- infrastructure
    |   |   |-- dbtests
    |   |   |-- dependencies
    |   |   `-- tests
    |   `-- repositories
    |       |-- basemodel
    |       |-- categories
    |       |-- items
    |       |-- products
    |       |-- providers
    |       |-- reviews
    |       `-- users
    |-- Dockerfile
    |-- go.mod
    |-- go.sum
    `-- main.go
|-- frontend
|   |-- app
|   |-- components
|   |-- hooks
|   |-- lib
|   |-- node_modules
|   |-- public
|   |-- styles
|   |-- components.json
|   |-- Dockerfile
|   |-- next.config.mjs
|   |-- next-env.d.ts
|   |-- package.json
|   |-- package-lock.json
|   |-- pnpm-lock.yaml
|   |-- postcss.config.mjs
|   |-- tailwind.config.ts
|   `-- tsconfig.json
|-- docker-compose.yml
`-- readme.md

```

