# migrate-mongo
Migrate mongodb collections by streaming way,include data and indexes. 

Powered by [MongoDB Go driver](https://github.com/mongodb/mongo-go-driver)

## This project is still on working, not finished.

## Features

* [x] Migrate mongo all data from db to db
* [x] Migrate mongo data from db to db by specified collections
* [x] Migrate mongo data include indexes, create indexes before inserting data
* [x] Batch inserting for migration
* [x] Interval between each single intersing

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
Migration options:

```
$ migrate-mongo help cln

Migrate specified collections

Usage:
  migrate-mongo cln <collections> [flags]

Flags:
      --all            Migrate all collections
  -b, --batch int32    Batch insert, count of each inserting (default 1)
  -d, --dst string     Destination mongodb uri (required) (default "mongodb://user:pwd@127.0.0.1/database2")
  -h, --help           help for cln
      --index          Include indexes, create indexes before inserting data
  -i, --interval int   Interval of each single insert, milliseconds
      --show-only      Only show details of source db collection, no migration operation
  -s, --src string     Source mongodb uri (required) (default "mongodb://user:pwd@127.0.0.1/database1")
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

`migrate-mongo cln test test2 --src mongodb://u:p@127.0.0.1:27017/db1 --dst mongodb://u:p@127.0.0.1:27017/db2`

outputs:

```
Collection details: [test test2]

Name: test
Count: 4
Indexes:
   _id_, userid_-1

Name: test2
Count: 3
Indexes:
   _id_

Start migration:

Done: test 4/4, elapsed: 159ms
Done: test2 3/3, elapsed: 112ms
```

## Test

TODO

## Tips

* You should avoid db writing while migrating

## LICENSE

[LICENCE MIT](https://github.com/Dongss/migrate-mongo/blob/master/LICENSE)
