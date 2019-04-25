# migrate-mongo
Migrate mongodb collections by streaming way,include data and indexes. 

Powered by [MongoDB Go driver](https://github.com/mongodb/mongo-go-driver)

## This project is still on working, not finished.

## Features

* [x] Migrate mongo all data from db to db
* [x] Migrate mongo data from db to db by specified collections
* [ ] Migrate mongo data include indexes
* [ ] Batch insert for migration
* [x] Interval between each single intersing for DB load

## Install

`go get -u github.com/Dongss/migrate-mongo`

## Usage

```
$ migrate-mongo --help

A Tool for data migrations between mongodb databases.
Migrations are by streaming way.
Complete documentation is available at https://github.com/Dongss/migrate-mongo

Usage:
  migrate-mongo [command]

Available Commands:
  cln         Migrate specified collections
  help        Help about any command
  version     Print the version of migrate-mongo

Flags:
  -h, --help   help for migrate-mongo

Use "migrate-mongo [command] --help" for more information about a command.
```

## Example

Overview:

`migrate-mongo cln test log --src mongodb://u:p@127.0.0.1:27017/db1 --dst mongodb://u:p@127.0.0.1:27017/db2 --show-only`

outputs:

```
Collection details: [test log]

Name: test
Count: 4
Indexes:
   _id_, userid_-1

Name: log
Count: 1548
Indexes:
   _id_
```

Migrate:

`
`migrate-mongo cln test --src mongodb://u:p@127.0.0.1:27017/db1 --dst mongodb://u:p@127.0.0.1:27017/db2`

outputs:

```
Collection details: [test]

Name: test
Count: 4
Indexes:
   _id_, userid_-1

Start migration:

Done: test 4/4, elapsed: 159ms
```

## Test

TODO

## Tips

* You should avoid db writing while migrating

## LICENSE

[LICENCE MIT](https://github.com/Dongss/migrate-mongo/blob/master/LICENSE)
