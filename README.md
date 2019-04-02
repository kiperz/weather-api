# weather-api

# Usage

Project assumes user installed docker+docker-compose.

Go to project directory then run
```docker-compose up```
app will compile and start on port 8080.

Routes available to user:
* POST /v1/api/bookmark/ (body should be json: {"name":"New York"})
* GET /v1/api/bookmark/
* POST /v1/api/query/ (body should be json: {"name":"New York"})
* GET /va/api/query?name=
