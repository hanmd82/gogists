## User Authentication

Goals
- How to implement basic signup, login and logout functionality for users
- Encrypt and store user passwords securely in database using bcrypt
- Verify that a user is logged in using middleware and sessions
- How to prevent Cross-Site Request Forgery (CSRF) attacks

---

### Create a Users Model

Connect to PostgreSQL DB and run the following SQL commands:

```sql
CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

\d users
/*
                                           Table "public.users"
     Column      |            Type             | Collation | Nullable |              Default
-----------------+-----------------------------+-----------+----------+-----------------------------------
 id              | integer                     |           | not null | nextval('users_id_seq'::regclass)
 name            | character varying(255)      |           | not null |
 email           | character varying(255)      |           | not null |
 hashed_password | character(60)               |           | not null |
 created_at      | timestamp without time zone |           | not null |
 active          | boolean                     |           | not null | true
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
    "users_uc_email" UNIQUE CONSTRAINT, btree (email)
*/
```
