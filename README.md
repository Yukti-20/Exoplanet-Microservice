# Exoplanet-Microservice

Here are the steps to execute the Exoplanet-Microservices

#### 1. Navigate to the project directory

```
    cd /path/to/Exoplanet-Microservice
```

#### 2. Initialize go module

```
    go mod init Exoplanet-Microservice
```

#### 3. Run module dependencies

```
    go mod tidy
```

#### 4. Build the project

```
    go build .
```

#### 5. Run the Project using Docker

```
   docker build -t exoplanet-microservice .
   docker run -p 8080:8080 exoplanet-microservice
```

#### 6. Run the below apis and can change the payload with the request type

```
    POST   http://localhost:8080/add/exoplanets
    GET    http://localhost:8080/list/exoplanets
    GET    http://localhost:8080/get/exoplanet/{id}
    POST   http://localhost:8080/update/exoplanet/{id}
    DELETE http://localhost:8080/delete/exoplanet/{id}
    GET    http://localhost:8080/fuel/estimation/{id}
```
