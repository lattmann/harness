# Practice API Implementation

As a practice and minimum working example, implement the following APIs and allow it only if the user logged in:

- Create a workspace: `POST /api/v1/split/api/v2/workspaces/`
- Get a workspace: `GET /api/v1/split/api/v2/workspaces/:id`
- Delete a workspace: `DELETE /api/v1/split/api/v2/workspaces/:id`

## Code Structure

### Database

The database has two different drivers:

- `postgres`
- `sqlite`

The `postgres` driver is the production driver and is used by default. The `sqlite` driver is used for testing and development.

In order to create new tables for the database you must create separate migration files for each backend.

Tables are implemented partially based on split.io's API documentation: [https://docs.split.io/reference/introduction](https://docs.split.io/reference/introduction)

[`app/store/database/migrate/postgres`](app/store/database/migrate/postgres)
- [`app/store/database/migrate/postgres/0097_create_extension_uuid.up.sql`](app/store/database/migrate/postgres/0097_create_extension_uuid.up.sql) - Create the uuid extension if it doesn't exist
- [`app/store/database/migrate/postgres/0098_create_table_split_workspaces.down.sql`](app/store/database/migrate/postgres/0098_create_table_split_workspaces.down.sql) - Drop the table if it exists
- [`app/store/database/migrate/postgres/0098_create_table_split_workspaces.up.sql`](app/store/database/migrate/postgres/0098_create_table_split_workspaces.up.sql) - Create the table - uses `UUID` whereas and harness tables use `SERIAL` id
- [`app/store/database/migrate/postgres/0099_create_table_split_environments.down.sql`](app/store/database/migrate/postgres/0099_create_table_split_environments.down.sql)
- [`app/store/database/migrate/postgres/0099_create_table_split_environments.up.sql`](app/store/database/migrate/postgres/0099_create_table_split_environments.up.sql)
- [`app/store/database/migrate/postgres/0100_create_table_split_traffic_types.down.sql`](app/store/database/migrate/postgres/0100_create_table_split_traffic_types.down.sql)
- [`app/store/database/migrate/postgres/0100_create_table_split_traffic_types.up.sql`](app/store/database/migrate/postgres/0100_create_table_split_traffic_types.up.sql)
- [`app/store/database/migrate/postgres/0101_create_table_split_segments.down.sql`](app/store/database/migrate/postgres/0101_create_table_split_segments.down.sql)
- [`app/store/database/migrate/postgres/0101_create_table_split_segments.up.sql`](app/store/database/migrate/postgres/0101_create_table_split_segments.up.sql)

`app/store/database/migrate/sqlite`: Not implemented

Rebuild the application and then restart it. Migration will take place automatically.

```bash
$ make build
$ ./gitness server .local.env
```

#### Type definitions for serialized data

[`types/split.go`](types/split.go)

#### Create databse stores

_**Function**: Defines the database store and provides the column name mappings for the database and any interactions with the database._

- Create a new store implementation `SplitWorkspaceStore` [`app/store/database/split_workspace.go`](app/store/database/split_workspace.go) - also provides the column name mappings for the database and any interactions with the database
- Create a store provider and add it to the wire set [`app/store/database/wire.go`](app/store/database/wire.go)

### API

#### Controller

_**Function**: Receives session and input data, passes request to the database store, and returns the result._

Create a new controller for the `split` API. We will use a single controller for all endpoints and data types. Note: Harness has separate controllers for each data type.

- Create a new folder `app/api/controller/split`
- New WireSet setup [`app/api/controller/split/wire.go`](app/api/controller/split/wire.go)
- Controller implementation [`app/api/controller/split/controller.go`](app/api/controller/split/controller.go)
- Provide the new controller in the wire set [`app/api/controller/wire.go`](app/api/controller/wire.go)

#### Handler

_**Function**: Receives session and input data, passes request to the controller, and returns the result._

Implement the handler for the `split` API.

Note: Harness uses one handler per endpoint.

- Create a new folder `app/api/handler/split`
- Handler implementation [`app/api/handler/split/create_workspace.go`](app/api/handler/split/create_workspace.go) - Implements both POST and GET -
- Add any query or path parsing logic into the request package [`app/api/request/split.go`](app/api/request/split.go) - e.g., `GetUUIDParam`

_TODO: Separate the handlers into different files_

#### Attach controller and handler to the API

- [`app/router/api.go`](app/router/api.go) - list your API endpoints here, load the controller and attach handlers


### Regenerate `cmd/wire_gen.go`

```bash
$ pushd cmd
$ go generate .
$ popd
```

## Rebuild, test, and run the application

```bash
$ make build
$ make test
$ ./gitness server .local.env
```
