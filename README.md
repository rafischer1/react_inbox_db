# GO Backend for React-Inbox 🍽

> CRUD routing for a RESTful API architecture in Go. PostgreSQL migrations and seeds managed with the vanilla postgres commands.

[Deployed DB:](https://fischer-go-inbox.herokuapp.com/)

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

## routes 🚍

  `/messages/{id}`

* mock `POST` request: http POST :3003/messages read=true starred=false selected=false subject='Subject Entry' body='Body Entry' labels='{"personal", "dev"}'

* mock `PUT` request: http PUT :3003/messages/1 read=true starred=false subject="Edited Subject" body="Edited Body" labels='{"personal"}'

## database migrations with *Goose*

[Goose 💸](https://github.com/pressly/goose)

`goose create messages`

`goose up` / `goose down` / `goose redo`
