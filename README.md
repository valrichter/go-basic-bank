# <img id="go" src="https://devicon-website.vercel.app/api/go/plain.svg?color=%2300ACD7" width="30" /> Proyecto extenso de Backend con Go 

Es un sistema basico del funcionamiento de un banco echo en Go. Aplicando distintos conceptos de Backend + CI/CD + AWS.
La idea es cubrir las operaciones basicas de CRUD y la tranferencia de dinero entre usuarios de la app
## 🔨 Tecnologias usadas:
- **Go**: go1.21.1 linux/amd64
- **PostgreSQL**: docker image postgres:15.4
- **Docker**: docker v24.0.6
- **CI**: GitHub Actions
- **SQLC**: sqlc-dev/sqlc v1.21.0
- **Migrate**: golang-migrate v4.16.2
- **Make**: GNU Make v4.3

### 📦 Herramietas:
- **Gin**: gin-gonic/gin v1.9.1
- **Testify**: stretchr/testify v1.8.4
- **Viper**: spf13/viper v1.16.0
- **GoMock**: golang/mock v1.6.0
- **JWT**: golang-jwt/jwt v3.2.2+incompatible
- **Paseto**: o1egl/paseto v1.0.0
## ⚡ Acciones realizadas durante el proyecto:

### 🗃️ Coneccion con la DB [PostgreSQL + sqlc]

**1.** Esquema de la DB y realacion entre tablas:
   - Crear una Account (Owner, Balance, Currency)
   - Registrar todos los cambios de balance en la cuenta (Entry)
   - Hacer transferencias de dinero entre 2 Accounts (Transfer)
   <img src="https://github.com/valrichter/basic-system-bank/assets/67121197/f0087f1e-ab3b-4532-a7bc-1a578c7c1e2c"/>

**2.** Configuracion de imagen de PostgreSQL en Docker y creacion de la DB mediante un archivo *.sql*:
   - Se agregaron indices (indexes) de los atributos mas importantes de cada tabla para mayor eficiencia a la hora de la busqueda

**3.** Creacion de versiones de la DB. Configuracion golang-migrate para hacer migraciones
s de la DB de una version a otra:
   - Se agrego un Makefile para mayor comodidad a la hora de ejecutar comandos necesarios
<img src="https://github.com/valrichter/basic-system-bank/assets/67121197/707a01c9-699c-427c-8838-16b422b891d0"/>

**4.** Generacion de CRUD basico para las tablas Account, Entry & Transfer con sqlc. Configuracion de sqlc para hacer consultas SQL con codigo Go:
   - Como funciona? Consulta SQL -> [sqlc] -> Codigo Go con interfaces para poder interactuar

**5.** Generacion de datos falsos y creacion de Unit Tests para CRUD de la tabla Account:
   - Utlizacion del archivo ```random.go```

**6.** Creacion de una transaccion ```StoreTx.go``` con las propiedades ACID para la transferencia de dinero entre 2 Accounts y su respectivo Unit Test:

<img src="https://github.com/valrichter/basic-system-bank/assets/67121197/4e3b1cf6-f593-46b7-a101-5a2e32f992b9"/>

**7.** Creacion de Unit Tests (TDD) con go routines para simular tracciones concurrentes y evitar transaction locks.
   
**8.** Modificacion del codigo para evitar situaciones deadlock y Unit Test para transactions deadlocks.

**9.** Estudio de los distintos Insolation Levels en PostgreSQL:
| Read Phenomena / Isonlation Levels ANSI | Read Uncommited | Read Commited | Repeatable Read | Serializable |
| :-------------------------------------: | :-------------: | :-----------: | :-------------: | :----------: |
|               Dirty Read                |       NO        |      NO       |       NO        |      NO      |
|           Non-Repeatable Read           |       SI        |      SI       |       NO        |      NO      |
|              Phantom Read               |       SI        |      SI       |       NO        |      NO      |
|          Serialization Anomaly          |       SI        |      SI       |       SI        |      NO      |

**10.**  Implementacion de Continouous Integration (CI) con GitHub Actions para garatizar la calidad del codigo y reducir posibles errores:

<img src="https://github.com/valrichter/basic-system-bank/assets/67121197/d7ac2106-9628-41db-a203-3e653bf30ddc"/>

***

### 🧩 Construccion de una RESTful HTTP JSON API [Gin + JWT + PASETO]

**1.** Implementacion de una RESTful HTTP API basico con Gin, configurancion del server y agregado de las funciones createAccount, getAccount by id y listAccount para listar cuentas mediante paginacion

**2.** Creacion de variables de entorno con ```.env``` y viper

**3.** Implementacion de Mock DB con GoMock para testear los metodos de ```account.go``` de la API HTTP y logrando en covertura del 100% del metodo GetAccounts

**4.** Implementacion de ```transfer.go``` de la API HTTP para enviar dinero entre dos cuentas y se agrego un ```validator.go``` para validar la Currency de las cuentas relacionadas con la transferencia de dinero

**5.** Se actualizao la base de datos y se agrego la tabla User para que cada usuario pueda tener distintas Accounts con diferentes Currency como ARS, UDS o EUR

<img src="https://github.com/valrichter/go-basic-bank/assets/67121197/54005e6b-ebad-4689-af1d-d1b602b25c9a"/>

**6.** Implentacion y test de CRUD users, manejo de errores de la db y fix de la API ```account.go``` para que funcione con la nueva tabla Users

**7.** Implentacion del la API ```user.go``` y encriptacion de la passwaord de los Users utilizando bcrypt para evitar "Rainbow Attacks"

**8.** Creacion de test para la funcion createUser de la API ```user.go```

**9.** Entendimiento de las defirencias entre JWT y PASETO. Problemas de seguridad de JWT y como funciona PASETO

**10.** Implentacion de generatcion de tokes de JWT y PASETO con sus test para la generacion de tokens

**11.** Implentacion de Paseto para login de Users

**12.** Implementacion de middleware de autenticación y reglas de autorización usando Gin

***

### ☁️ Deployar la aplicacion a produccion [Docker + Kubernetes + AWS]

**1.** Se creo un archivo Dockerfile multietapa para crear la imagen con la app que contenga solo el binario ejecutable

**2.** Se conecto los dos conatiners, el que contiene la db y el que contiene el binario ejecutable, a una misma network para que puedan comunicarse entre ellos

**3.** Configuracion de docker-compose para inicializar los dos servicios y coordinarlos

**4.** Investigacion de como usar AWS para conrrer los conatiners de docker en la nube

**5.** Configuracion de GitHub Actions para poder automatizar el deploy de un container en ECR de AWS


