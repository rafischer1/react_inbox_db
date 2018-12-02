# GO Backend for React-Inbox 🍽

> CRUD routing for a RESTful API architecture in Go. PostgreSQL migrations and seeds managed with the Goose package import.

```haskell
CREATE TABLE messages (
 id serial PRIMARY KEY,
 read boolean,
 starred boolean,
 selected boolean,
 subject VARCHAR (50) UNIQUE NOT NULL,
 body VARCHAR (255) UNIQUE NOT NULL,
 labels VARCHAR (255) UNIQUE NOT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL
);
```

### routes

  `/messages`
