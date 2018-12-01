# GO Backend for React-Inbox 🍽

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
