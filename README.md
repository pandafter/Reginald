# Reginald

Sistema modular de gestión con backend en microservicios y frontend como interfaz de administración.

![RegiPresentation](Assets/RegiPresentation.png)

---

## Tabla de contenidos

- [Requisitos](#requisitos)
- [Descarga e instalación](#descarga-e-instalación)
- [Backend (Docker)](#backend-docker)
- [Frontend (gestor)](#frontend-gestor)
- [Assets](#assets)

---

## Requisitos

| Componente | Requisito |
|------------|-----------|
| **Backend** | Docker y Docker Compose |
| **Frontend** | Node.js ≥20.19.0 o ≥22.12.0, npm |

---

## Descarga e instalación

```bash
git clone https://github.com/<tu-usuario>/Reginald.git
cd Reginald
```

---

## Backend (Docker)

El backend es una arquitectura de microservicios en Go con:
- **Gateway** (puerto 8088) – API unificada
- **Auth** – autenticación y JWT
- **Users** – gestión de usuarios
- **Tasks** – gestión de tareas
- **Requests** – gestión de solicitudes
- **PostgreSQL** – base de datos

### 1. Entrar al directorio del backend

```bash
cd reginald_backend
```

### 2. Configurar variables de entorno

Copiar el archivo de ejemplo y ajustar los valores:

```bash
cp .example.env .env
```

Editar `.env` con tus valores (por ejemplo):

```env
# Admin inicial
ADMIN_NAME=AdminExample
ADMIN_EMAIL=admin@reginald.com
ADMIN_PASSWORD=admin123

# Base de datos
DB_USER=postgres
DB_PASSWORD=mypassword
DB_NAME=reginald
DB_PORT=5432

# Seguridad (¡cambiar en producción!)
JWT_SECRET=super-secret-key-12345
```

### 3. Levantar los servicios con Docker Compose

```bash
docker compose up -d
```

### 4. Verificar que el backend responde

- **API Gateway:** `http://localhost:8088`
- **Health check:** `http://localhost:8088/health`

### 5. Detener los servicios

```bash
docker compose down
```

---

## Frontend (gestor)

El frontend es la interfaz web (Vue 3 + Vite) que gestiona el backend. Se conecta al Gateway en el puerto **8088**.

### 1. Entrar al directorio del frontend

```bash
cd reginald_frontend
```

### 2. Instalar dependencias

```bash
npm install
```

### 3. Ejecutar en modo desarrollo

```bash
npm run dev
```

El frontend se sirve en `http://localhost:5173` y hace proxy de `/api` al backend en `http://127.0.0.1:8088`. El backend debe estar corriendo.

### 4. Producción

```bash
npm run build
npm run preview
```

Para producción, sirve los archivos generados en `dist/` y configura el reverse proxy para que `/api` apunte a tu Gateway.

---

## Assets

| Imagen | Descripción |
|--------|-------------|
| [RegiPresentation.png](Assets/RegiPresentation.png) | Ilustración explicativa del software |
| [screenShot1.png](Assets/screenShot1.png) | Captura de pantalla |
| [screenShot2.png](Assets/screenShot2.png) | Captura de pantalla |
| [screenShot3.png](Assets/screenShot3.png) | Captura de pantalla |
| [screenShot4.png](Assets/screenShot4.png) | Captura de pantalla |

---

## Resumen de puertos

| Servicio | Puerto |
|----------|--------|
| Frontend (dev) | 5173 |
| Gateway (API) | 8088 |
| PostgreSQL | 5432 |
