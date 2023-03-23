# BPJS Test

This project is using PostgreSQL for the database and i've created auto migration inside `main.go`. Please create a database with the name "my_db" first for easy running. This project has 1 route, service, and database only to accomplised the test task. Also i've provide a screenshot about the DB and Postman each view.

#### Route
### GET /person
#### Request
```javascript
[
   {
      "id": 1,
      "customer": "John Smith 1",
      "quantity": 10,
      "price": 100,
      "timestamp": "2023-03-23T11:30:06+07:00"
   }, ...
]
```

#### Response
```javascript
{
    "error": false,
    "message": "success insert data to DB, time taken 114ms"
}
```

#### Screenshot
![alt text](https://github.com/daudfauzy98/bpjs-test/blob/main/psql.jpg)
![alt text](https://github.com/daudfauzy98/bpjs-test/blob/main/postman.png)
