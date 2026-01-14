# sample-client
a simple go client for the simple web-server.

## installation

- install go 
- clone the repo
- install dependencies

## to implement

- [x] base route : `/`

- [x] insert a user to the database : `/register`
 - provide : `firstName`, `lastName` `email` `password`
    `phone` `age` `job` in the request body
 - the email is unique
 - all fields Not nullable except (age, phone, job)

- [x] get all users in the database : `/users`
 - send an email (must be a user email)

- [ ] login and get a jwt token: `/auth`.
 - send a valid credentials {email , password} in the request body
 - recive a token if its valid cedentials.

- [ ] get a specific user info : `/query/{email}`
 - send an email parameter the route
 - send a valid token in the `Authorization` request header

## future implementation

- [ ] beatify output
- [ ] make a UI (web page)
