# Gin API Example

An example project layout for Go applications written with the [Gin](https://github.com/gin-gonic/gin) framework.
I'm also using [GORM](https://github.com/go-gorm/gorm) as an ORM and [golang-migrate](https://github.com/golang-migrate/migrate)
to handle database migrations.

This template / example is an ongoing work in progress, so depending on the timing there may be bugs.

I've borrowed heavily from [@benbjohnson/standard-package-layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) and [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

# Prerequistes

* [golang-migrate](https://github.com/golang-migrate/migrate) 
  * `brew install golang-migrate`

# Try It Out

```
$ docker-compose up
$ make db-up
$ make seed
$ make run
```

* API docs are available at http://localhost:8081
