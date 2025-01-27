# Test task - REST API for finance operations

<h3>Configuration</h3>

```
sudo docker run --name=bank -e POSTGRES_PASSWORD=54321 -p 5422:5432 -d postgres:16.2
```

<h4> ENV Example </h4>

```
DB_URL=postgres://postgres:54321@localhost:5422/postgres?sslmode=disable

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:54321@localhost:5422/postgres?sslmode=disable
GOOSE_MIGRATION_DIR=./migrations
```

<h4>How to run</h4>

<b>To build docker-compose</b>

```
make run
```

<b>To make executable</b>

```
make app 
```

<h3>API EndPoints</h3>

<b>1. Increase user's balance</b>

```
POST /api/v1/up
```

<b>Body request</b>

```
{
  "user_id": 1,
  "amount": 100.50
}
```

<b>Params</b>

<ul>
   <li> user_id (int): user's ID that need to increase. </li>
   <li> amount (float64): amount by which the balance should be increased</li>
</ul>


<b>Example response</b>

```
{
  "status": "success"
}

{
  "error": "invalid user ID"
}
```

<b>2. Transferring money between users</b>

```
POST /api/v1/transfer
```

<b> Body request </b>

```
{
  "from_user_id": 1,
  "to_user_id": 2,
  "sender_name": "Alice",
  "receiver_name": "Bob",
  "amount": 50.75
}
```

<b>Params</b>

<ul>
    <li>from_user_id (int): sender's ID.</li>
    <li>to_user_id (int): receiver's ID.</li>
    <li>sender_name (string): sender's name.</li>
    <li>receiver_name (string): receiver's name.</li>
    <li>amount (float64): amount of transfer.</li>
</ul>

<b>Example response</b>

```
{
  "Money has transfered to ": "Bob"
}

{
  "error": "any error"
}
```

<b>3. Get last user's txs</b>

```
GET /api/v1/tx/{userID}
```

<ul>
   <li>user's ID who need to know his last txs </li>
</ul>

<b>Example response</b>

```
[
    {
        "Amount": 75,
        "CreatedAt": "2023-10-01T14:00:00Z",
        "Receiver_name": "Charlie",
        "Sender_name": "Alice",
        "Type": "transfer"
    },
    {
        "Amount": 200,
        "CreatedAt": "2023-10-01T10:00:00Z",
        "Receiver_name": "Bob",
        "Sender_name": "Alice",
        "Type": "transfer"
    }
]

{
  "error": "invalid user ID"
}
```
