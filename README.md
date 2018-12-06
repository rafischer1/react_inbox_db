# GO Backend for React-Inbox 🍽

> CRUD routing for a RESTful API architecture in Go. PostgreSQL migrations and seeds managed with the vanilla postgres commands.

```haskell
CREATE TABLE messages (
 id serial PRIMARY KEY,
 read boolean,
 starred boolean,
 selected boolean,
 subject VARCHAR (255) UNIQUE NOT NULL,
 body VARCHAR (255) UNIQUE NOT NULL,
 labels text[],
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NOT NULL
);
```

## routes

  `/messages/{id}`

* mock post request: read=true starred=false selected=false subject="Test Entry" body="Test Entry Body" labels="{"personal", "dev"}"

## database creation with *Goose*

[Goose 💸](https://github.com/pressly/goose)

`goose create messages`

`goose up` / `goose down` / `goose redo`
