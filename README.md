## POC - GOLANG VALIDATE: ARCH, DEPENDENCIES AND DIRECTORIES!

> **THIS IS JUST A POC** - To validate the architecture, dependencies and folder structure of a golang system.

This POC uses the [arch-go](https://github.com/fdaines/arch-go) *(to validate architecture and dependencies - file: [arch-go.yml](arch-go.yml))* and [directory-validator](https://github.com/goerwin/directory-validator) *(to validate the directories structure - file: [.directoryvalidator.json](.directoryvalidator.json))* tools.

**Note:** Disregard the code, it was created only to validate the structure, not the code!

----

### Run

#### Requirements
- [Docker](https://docs.docker.com/engine/install/) >= 15;
- [Golang](https://go.dev/doc/install) >= 1.16.x;

#### Running: Validations

**Install tools: REQUIRED**

```shell
make internal-tools-install
```

**Validate: Architecture and Dependencies**

```shell
make lint-arch
```

**Validate: Directories Structure**

```shell
make lint-directories
```

**Validate: ALL**

```shell
make lint
```

----


### Imports and dependencies not allowed:

#### Internal
- `pkg/infra` cannot import
  - `pkg/domain/*`;
- `pkg/domain/$domain_module/entrypoint/*` cannot imports:
  - `pkg/domain/$domain_module/entrypoint/*` *(other)*;
  - `pkg/domain/$domain_module/service/*`;
  - `pkg/domain/$domain_module/repository/*`;
- `pkg/domain/$domain_module/usecase/*` cannot imports:
  - `pkg/domain/$domain_module/entrypoint/*`;
  - `pkg/domain/$domain_module/repository/*`;
- `pkg/domain/$domain_module/service/*` cannot imports:
  - `pkg/domain/$domain_module/entrypoint/*`;
  - `pkg/domain/$domain_module/usecase/*`;
- `pkg/domain/$domain_module/repository/*` cannot imports:
  - `pkg/domain/$domain_module/repository/*` *(other)*;
  - `pkg/domain/$domain_module/entrypoint/*`;
  - `pkg/domain/$domain_module/service/*`;

#### External
- `pkg/domain/$domain_module/entrypoint/*` cannot imports:
    - `.*bigquery.*`
    - `.*cache.*`
    - `.*fx.*`
    - `.*gorm.*`
    - `.*kvs.*`
    - `.*mysql.*`
    - `.*postgres.*`
    - `.*sql.*`
- `pkg/domain/$domain_module/usecase/*` cannot imports:
    - `.*api.*`
    - `.*bigquery.*`
    - `.*cache.*`
    - `.*chi.*`
    - `.*fx.*`
    - `.*gin.*`
    - `.*gorilla.*`
    - `.*gorm.*`
    - `.*grpc.*`
    - `.*header.*`
    - `.*http.*`
    - `.*json.*`
    - `.*jwt.*`
    - `.*kvs.*`
    - `.*mux.*`
    - `.*mysql.*`
    - `.*postgres.*`
    - `.*protobuf.*`
    - `.*querystring.*`
    - `.*sql.*`
    - `.*websocket.*`
- `pkg/domain/$domain_module/service/*` cannot imports:
    - `.*api.*`
    - `.*bigquery.*`
    - `.*cache.*`
    - `.*chi.*`
    - `.*fx.*`
    - `.*gin.*`
    - `.*gorilla.*`
    - `.*gorm.*`
    - `.*grpc.*`
    - `.*header.*`
    - `.*http.*`
    - `.*json.*`
    - `.*jwt.*`
    - `.*kvs.*`
    - `.*mux.*`
    - `.*mysql.*`
    - `.*postgres.*`
    - `.*protobuf.*`
    - `.*querystring.*`
    - `.*sql.*`
    - `.*websocket.*`
- `pkg/domain/$domain_module/repository/http/*` cannot imports:
    - `.*fx.*`
    - `.*gorm.*`
    - `.*mysql.*`
    - `.*postgres.*`
    - `.*sql.*`
- `pkg/domain/$domain_module/repository/sql/*` cannot imports:
    - `.*fx.*`
    - `.*bigquery.*`
    - `.*cache.*`
    - `.*kvs.*`
    - `.*api.*`
    - `.*chi.*`
    - `.*gin.*`
    - `.*gorilla.*`
    - `.*grpc.*`
    - `.*header.*`
    - `.*http.*`
    - `.*jwt.*`
    - `.*mux.*`
    - `.*protobuf.*`
    - `.*querystring.*`
    - `.*websocket.*`

#### Limitations
This validates the direct import, it is not validating cases of dependency injection or *duck typing* outside: `pkg/domain/$domain_module`.

### Structure allowed:

```
.
└── pkg
    ├── domain
    │   └── $domain_module (there can be N* modules)
    │       ├── entity.go
    │       ├── entity_test.go
    │       ├── entrypoint
    │       │   └── rest
    │       │       ├── mocks
    │       │       │   └── $mocks.go (there can be N* mocks)
    │       │       ├── rest.go
    │       │       └── rest_test.go
    │       ├── repository
    │       │   ├── http
    │       │   │   ├── http.go
    │       │   │   └── http_test.go
    │       │   └── sql
    │       │       ├── sql.go
    │       │       └── sql_test.go
    │       ├── service
    │       │   ├── mocks
    │       │   │   └── $mocks.go (there can be N* mocks)
    │       │   ├── service.go
    │       │   └── service_test.go
    │       └── usecase
    │           ├── mocks
    │           │   └── $mocks.go (there can be N* mocks)
    │           ├── $uc.go
    │           └── $uc_test.go
    └── infra
        └── $generic_module (there can be N* generic module)
            ├── $generic_module.go
            └── $generic_module_test.go
```

#### What is not allowed

- Create a folder that is not `domain` or `infra` in `pkg`;
- Create files in `pkg`;
- Create a folder not being `entrypoint`, `usecase`, `service` or `repository` in `pkg/domain/$domain_module`;
- Create a file that is not `entity.go` or `entity_test.go` in `pkg/domain/$domain_module`;
- Create a folder that is not `rest` in `pkg/domain/$domain_module/entrypoint`;
- Create files in `pkg/domain/$domain_module/entrypoint`;
- Create a folder that is not `mocks` in `pkg/domain/$domain_module/entrypoint/rest`;
- Create a file that is not `rest.go` or `rest_test.go` in `pkg/domain/$domain_module/entrypoint/rest`;
- Create a folder that is not `mocks` in `pkg/domain/$domain_module/usecase`;
- Create a folder that is not `mocks` in `pkg/domain/$domain_module/service`;
- Create a file that is not `service.go` or `service_test.go` in `pkg/domain/$domain_module/service`;
- Create a folder not being `http`, or `sql` in `pkg/domain/$domain_module/repository`;
- Create files in `pkg/domain/$domain_module/repository`;
- Create a folder that is not `mocks` in `pkg/domain/$domain_module/repository/http`;
- Create a file that is not `http.go` or `http_test.go` in `pkg/domain/$domain_module/repository/http`;
- Create a folder that is not `mocks` in `pkg/domain/$domain_module/repository/sql`;
- Create a file that is not `sql.go` or `sql_test.go` in `pkg/domain/$domain_module/repository/sql`;

**Example**: project structure that follows this pattern:

```
.
└── pkg
    ├── domain
    │   ├── domain01
    │   │   ├── entity.go
    │   │   ├── entrypoint
    │   │   │   └── rest
    │   │   │       ├── mocks
    │   │   │       │   └── service.go
    │   │   │       ├── rest.go
    │   │   │       └── rest_test.go
    │   │   ├── repository
    │   │   │   ├── http
    │   │   │   │   ├── http.go
    │   │   │   │   └── http_test.go
    │   │   │   └── sql
    │   │   │       ├── sql.go
    │   │   │       └── sql_test.go
    │   │   ├── service
    │   │   │   ├── mocks
    │   │   │   │   └── repository.go
    │   │   │   ├── service.go
    │   │   │   └── service_test.go
    │   │   └── usecase
    │   │       ├── delete.go
    │   │       ├── delete_test.go
    │   │       ├── find.go
    │   │       ├── find_test.go
    │   │       ├── mocks
    │   │       │   ├── serviceDelete.go
    │   │       │   ├── serviceFind.go
    │   │       │   └── serviceSave.go
    │   │       ├── save.go
    │   │       └── save_test.go
    │   ├── domain02
    │   │   ├── entity.go
    │   │   ├── repository
    │   │   │   └── sql
    │   │   │       ├── sql.go
    │   │   │       └── sql_test.go
    │   │   └── service
    │   │       ├── mocks
    │   │       │   └── repository.go
    │   │       ├── service.go
    │   │       └── service_test.go
    │   └── domain03
    │       ├── entity.go
    │       ├── entrypoint
    │       │   └── rest
    │       │       ├── mocks
    │       │       │   └── service.go
    │       │       ├── rest.go
    │       │       └── rest_test.go
    │       ├── repository
    │       │   └── sql
    │       │       ├── sql.go
    │       │       └── sql_test.go
    │       ├── service
    │       │   ├── mocks
    │       │   │   └── repository.go
    │       │   ├── service.go
    │       │   └── service_test.go
    │       └── usecase
    │           ├── mocks
    │           │   ├── serviceDomain02Save.go
    │           │   └── serviceSave.go
    │           ├── save.go
    │           └── save_test.go
    └── infra
        ├── generic01
        │   ├── generic01.go
        │   └── generic01_test.go
        ├── generic02
        │   ├── generic02.go
        │   └── generic02_test.go
        └── generic03
            ├── generic03.go
            └── generic03_test.go
```
