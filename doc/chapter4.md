## Database-Driven Responses

### Setting up PostgreSQL
Create `gogists` database, and `gists` table with index on `created_at` field
```sql
CREATE DATABASE gogists ENCODING UTF8;
\c gogists
-- You are now connected to database "gogists" as user "mhan".

CREATE TABLE gists (
  id SERIAL NOT NULL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  expires_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_gists_created_at ON gists(created_at);

\d gists
/*
                                        Table "public.gists"
   Column   |            Type             | Collation | Nullable |              Default
------------+-----------------------------+-----------+----------+-----------------------------------
 id         | integer                     |           | not null | nextval('gists_id_seq'::regclass)
 title      | character varying(100)      |           | not null |
 content    | text                        |           | not null |
 created_at | timestamp without time zone |           | not null |
 expires_at | timestamp without time zone |           | not null |
Indexes:
    "gists_pkey" PRIMARY KEY, btree (id)
    "idx_gists_created_at" btree (created_at)
*/
```

Seed example data
```sql
INSERT INTO gists (title, content, created_at, expires_at) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    now() at time zone 'utc',
    now() at time zone 'utc' + 365 * INTERVAL '1 day'
);

INSERT INTO gists (title, content, created_at, expires_at) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    now() at time zone 'utc',
    now() at time zone 'utc' + 365 * INTERVAL '1 day'
);

INSERT INTO gists (title, content, created_at, expires_at) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    now() at time zone 'utc',
    now() at time zone 'utc' + 7 * INTERVAL '1 day'
);
```

Create new user `web` and grant access to `gists` table
```sql
CREATE USER web WITH PASSWORD '******';
GRANT SELECT,INSERT,UPDATE ON gists TO web;

\dp
/*
                                  Access privileges
 Schema |     Name     |   Type   | Access privileges | Column privileges | Policies
--------+--------------+----------+-------------------+-------------------+----------
 public | gists        | table    | mhan=arwdDxt/mhan+|                   |
        |              |          | web=arw/mhan      |                   |
 public | gists_id_seq | sequence |                   |                   |
*/
```

Verify permissions of user `web`
```sql
$ psql -U web -d gogists

SELECT id, title, expires_at FROM gists;
/*
 id |         title          |         expires_at
----+------------------------+----------------------------
  1 | An old silent pond     | 2021-07-22 14:37:34.155904
  2 | Over the wintry forest | 2021-07-22 14:37:43.446104
  3 | First autumn morning   | 2020-07-29 14:37:49.932896
(3 rows)
*/

DROP TABLE gists;
-- ERROR:  must be owner of table gists
```

---
### Installing a Database Driver

Reference: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

```
go get -u github.com/lib/pq
```

---
### Creating a Database Connection Pool

- `sql.Open()` doesn’t actually create any connections - it just initializes the connection pool for future use. Actual connections to the database are established lazily.
- use `db.Ping()` to create a connection and check for any errors.
