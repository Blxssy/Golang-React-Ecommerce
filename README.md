# Ecommerce

_This project is written in Go and uses the React for building frontend. A PostgreSQL database is used to store project's data._

## Config Variables

| Name                | Description                                                       |
|---------------------|-------------------------------------------------------------------|
| env: 'local'        | Defines the environment in which the application is running       |
| dialect: 'postgres' | Specifies the type of database the application is working with    |
| host: 'db'          | Indicates the host where the database is located                  |
| port: '5432'        | Specifies the connection port                                     |
| name: 'e-commerce'  | The name of the database                                          |
| username: 'postgres'| The username the application will use to connect to the database  |
| migration: true     | Specifies autmatic database migrations applying on app start      |


