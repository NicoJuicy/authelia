---
title: "authelia storage migrate up"
description: "Reference for the authelia storage migrate up command."
lead: ""
date: 2024-03-14T06:00:14+11:00
draft: false
images: []
weight: 905
toc: true
seo:
  title: "" # custom title (optional)
  description: "" # custom description (recommended)
  canonical: "" # custom canonical URL (optional)
  noindex: false # false (default) or true
---

## authelia storage migrate up

Perform a migration up

### Synopsis

Perform a migration up.

This subcommand performs the schema migrations available in this version of Authelia which are greater than the current
schema version of the database. By default this will migrate up to the latest available, but you can customize this.

```
authelia storage migrate up [flags]
```

### Examples

```
authelia storage migrate up
authelia storage migrate up --config config.yml
authelia storage migrate up --target 20 --config config.yml
authelia storage migrate up --encryption-key b3453fde-ecc2-4a1f-9422-2707ddbed495 --postgres.address tcp://postgres:5432 --postgres.password autheliapw
```

### Options

```
  -h, --help         help for up
  -t, --target int   sets the version to migrate to, by default this is the latest version
```

### Options inherited from parent commands

```
  -c, --config strings                        configuration files or directories to load, for more information run 'authelia -h authelia config' (default [configuration.yml])
      --config.experimental.filters strings   list of filters to apply to all configuration files, for more information run 'authelia -h authelia filters'
      --encryption-key string                 the storage encryption key to use
      --mysql.address string                  the MySQL server address (default "tcp://127.0.0.1:3306")
      --mysql.database string                 the MySQL database name (default "authelia")
      --mysql.password string                 the MySQL password
      --mysql.username string                 the MySQL username (default "authelia")
      --postgres.address string               the PostgreSQL server address (default "tcp://127.0.0.1:5432")
      --postgres.database string              the PostgreSQL database name (default "authelia")
      --postgres.password string              the PostgreSQL password
      --postgres.schema string                the PostgreSQL schema name (default "public")
      --postgres.username string              the PostgreSQL username (default "authelia")
      --sqlite.path string                    the SQLite database path
```

### SEE ALSO

* [authelia storage migrate](authelia_storage_migrate.md)	 - Perform or list migrations

