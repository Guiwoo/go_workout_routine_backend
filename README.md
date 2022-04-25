# WorkOut Routine Project

    - Go, GraphQL, PostgreSQL

## Why GraphQL ?

- [Click here Explain why graphql](howtographql.com/basics/1-graphql-is-the-better-rest/)

## Which GrpahQL Package are you gonna use ?

- 1. [Go grpahQL](https://tutorialedge.net/golang/go-graphql-beginners-tutorial/)
  - don't like it becaz not supporting the webpage

## PostgreSQL

- Need to study sql
  [Postgresql](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)
- Gorm vs Xorm
  - Gorm has 20k stars, Xorm has 6.3k stars
- Xorm why ? faster than gorm on the bench testing
  [Check here](https://sumit-agarwal.medium.com/gorm-vs-xorm-part-1-d156ba9de404)

## User

[x] create User "need to do hash the password and send if it's loggin on the header"

- [x] hashing passsword

[] login user

- [x] send back jwt token if it's currect information in the header
  - BUT NOT SURE HOW TO CHECK HEADER instead use rootValue
  - Need to Check another query
    - About header and rootValue
- [x] if jwt-token on header ? pass the query
