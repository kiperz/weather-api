# weather-api

# Usage

Project assumes user installed docker and docker-compose.

Setup your API_KEY in .env file

Go to project directory then run
```docker-compose up``` (it is required to run this command twice, at first time database is created and application can't connect)
app will compile and start on port 8080.

Routes available to user:
* POST /v1/api/bookmark/ (body should be json: {"name":"New York"})
* GET /v1/api/bookmark/
* POST /v1/api/query/ (body should be json: {"name":"New York"})
* GET /v1/api/query/stats?name=
