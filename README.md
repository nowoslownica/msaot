# msaot

## Build

* Install go
* Run ```go build```
* Build is ready - run `./msaot -h` binary

## Transcribe

```./msaot l2c <word>```

## Regenerate models

### Prerequisites

- ```go get -u -t github.com/volatiletech/sqlboiler```
- ```go get -u -t github.com/volatiletech/sqlboiler-sqlite3```

### Regeneration

- If the db structure changes, remove the db file (or add the logic of migration to 0000n+1 migration file)
- Go to migrations.go file
- Call go generate functions
- Run TestMigrateSchema from migrations_test.go
- Run sqlboiler sqlite3 from /db/sqlite folder