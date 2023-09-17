# Proyecto de Backend extenso en Go <img id="go" src="https://devicon-website.vercel.app/api/go/plain.svg?color=%2300ACD7" width="30" />

## 🔨 Herramietas usadas:
- GNU Make 4.3
- go version go1.21.0 linux/amd64
- golang-migrate 4.16.2
- Docker version 24.0.6
- Docker image postgres:15.4
- sqlc v1.20.0

## ⚡ Acciones realizadas durante el proyecto:

### Trabajando con DB [PostgreSQL + sqlc]

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

**5.** Generacion de datos falsos y creacion de Unit Tests para los CRUD de las tablas Account, Entry & Transfer:
   - Utlizacion del archivo ```random.go```

**6.** Creacion de una transaccion ```StoreTx.go``` con las propiedades ACID para la transferencia de dinero entre 2 Accounts y su respectivo Unit Test:

<img src="https://github.com/valrichter/basic-system-bank/assets/67121197/4e3b1cf6-f593-46b7-a101-5a2e32f992b9"/>

**7.** Creacion de Unit Tests (TDD) con go routines para simular tracciones concurrentes y evitar transaction locks.
   
**8.** Modificacion del codigo para evitar situaciones deadlock y Unit Test para transactions deadlocks.

**9.** Insolation Levels en PostgreSQL:
| Read Phenomena / Isonlation Levels ANSI | Read Uncommited | Read Commited | Repeatable Read | Serializable |
| :-------------------------------------: | :-------------: | :-----------: | :-------------: | :----------: |
| Dirty Read                              | NO              | NO            | NO              | NO           |
| Non-Repeatable Read                     | SI              | SI            | NO              | NO           |
| Phantom Read                            | SI              | SI            | NO              | NO           |
| Serialization Anomaly                   | SI              | SI            | SI              | NO           |

**10.**  Creacion de Continouous Integration (CI) con GitHub Actions para garatizar la calidad del codigo y reducir posibles errores:

<img src="https://github.com/valrichter/basic-system-bank/assets/67121197/d7ac2106-9628-41db-a203-3e653bf30ddc"/>

***

### Construccione de una RESTful HTTP JSON API [Gin + JWT + PASETO]
