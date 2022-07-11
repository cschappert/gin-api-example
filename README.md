# Gin API Example

An example project layout for Go applications written with the [Gin](https://github.com/gin-gonic/gin) framework.
I'm also using [GORM](https://github.com/go-gorm/gorm) as an ORM and [golang-migrate](https://github.com/golang-migrate/migrate)
to handle database migrations.

This template / example is an ongoing work in progress, so depending on the timing there may be bugs.

I've borrowed heavily from [@benbjohnson/standard-package-layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) and [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

# Prerequistes

* [golang-migrate](https://github.com/golang-migrate/migrate) 
  * `brew install golang-migrate`
* Mocks are generated with [mockery](https://github.com/vektra/mockery)
  * `brew install mockery`

# Try It Out

```
$ docker-compose up
$ make db-up
$ make seed
$ make run
```

* API docs are available at http://localhost:8081

# Run Unit Tests

```
$ make unit-test
```

# Generate Mocks

* For example, after adding a method to a repository interface

```
$ make mocks
```

# Project Structure

## Project Root

Contains the `docker-compose.yml`, `Makefile` for executing common tasks (test, run, etc.), `go.mod` file and so on.

## `api`

The API documentation. Is this case, it's an `openapi.yml` file.

## `cmd`

The go source code used to set up and run the application. Effectively the executable to run on the command line. You can either run `go run cmd/api/main.go` to start the application or `go run cmd/seed/main.go` to seed the DB with some test data.

## `db/migrations`

Database migrations go here. I'm using [golang-migrate](https://github.com/golang-migrate/migrate) to apply the migrations.

## `mocks`

Repository mocks are all in here. They are auto-generated using mockery (see the Makefile) and can be used to mock out the DB when unit testing.

## `scripts`

Any scripts you might have for your own convenience can go here.

## `pkg` (Top Level)

The Go source files here all implement application / business logic. It is effectively the "service layer". The source here basically does everything that does **not** depend on any knowledge re: the specifics of storing data to a DB, handling HTTP requests, etc.
If you take a look at the source, you will see some structs (used to represent business entities), some interfaces (that allow the service layer to
get what it needs from other, infrastructure focused layers) and some methods (that simply implement business logic and don't have any infrastructure
dependencies). You'll also see an accompanying `*_test.go` file that tests the logic here.

## `pkg/http`

The code here deals with all of the details of receiving requests, marshalling / unmarshalling JSON and returning responses. It is basically the
"handler" code.

## `pkg/storage/mysql`

The code here represents the "repository" layer and deals with the specifics of getting data into and out of a MySQL DB. The code here implements an
interface defined in the service layer. It can be passed in during application start-up to give the service layer a way of storing or retrieving data
(or not if you want to unit test with mocks). In addition to implementing the repository logic, the code here also defines some structs. This may seem
redundant as I've already defined my types in the service layer, but the types defined in the repository layer are defined to match perfectly to the DB
tables that they represent, whereas the types defined in the service layer are defined to work as well as possible in the service layer without being
bound to whatever the DB table structure may be.

I also like to keep my repository layer structs and service layer structs separate because I don't want something like a new column added to a table
leaking in to a service that doesn't have any use for that new column.

Lastly, the code to connect to the DB and get the ORM set up is here as well.