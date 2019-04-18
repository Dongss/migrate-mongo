# migrate-mongo
Migrate mongodb collections by streaming way,include data and indexes. 

Powered by [MongoDB Go driver](https://github.com/mongodb/mongo-go-driver)

## This project is still on working, not finished.

## Features

* Migrate mongo all data from db to db
* Migrate mongo data from db to db by specified collections
* Migrate mongo data include indexes

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
2019/04/17 18:38:42 Collection details: [test log]

Name: test
Count: 2
Indexes:
   _id_, userid_-1

Name: log
Count: 1544
Indexes:
   _id_
```

Migrate:
`
`migrate-mongo cln test --src mongodb://u:p@127.0.0.1:27017/db1 --dst mongodb://u:p@127.0.0.1:27017/db2`

outputs:

```
2019/04/18 10:50:16 Collection details: [test]

Name: test
Count: 2
Indexes:
   _id_, userid_-1

2019/04/18 10:50:17 Start migration:

start: test
done: test, count: 2, elapsed: 428ms
```

## Test

TODO

## LICENSE

[LICENCE MIT](https://github.com/Dongss/migrate-mongo/blob/master/LICENSE)
