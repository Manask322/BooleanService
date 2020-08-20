# BooleanService
Boolean service REST API with token Authentications 

#

### To run the application using docker
* change `DB_HOST` to `booleanservice-mysql`
* change `DB_PORT` to `33305`
* #### RUN `docker-compose up --remove-orphans`
* The Application is up and running at [localhost:8080](http://localhost:8080)   

### To run the application locally
* `cd` to the root directory
* #### RUN  `go build -o booleanservice .`
* #### RUN `./booleanservice`