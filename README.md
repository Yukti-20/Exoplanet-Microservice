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

<img width="1487" alt="Screenshot 2024-06-01 at 7 15 56 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/9a236c1b-e4b6-451e-b484-15a660933751">
<img width="1422" alt="Screenshot 2024-06-01 at 6 45 26 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/6b610544-a2eb-4779-a724-547b5c4ac6e3">
<img width="1426" alt="Screenshot 2024-06-01 at 6 41 22 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/6131da06-1fa9-47ce-b34a-5068055610ad">
<img width="1423" alt="Screenshot 2024-06-01 at 6 41 12 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/49672c45-56d6-428e-be30-900fd21d20da">
<img width="1424" alt="Screenshot 2024-06-01 at 6 40 13 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/18b949df-882f-40f5-aa7a-2250f398b6fe">
<img width="1423" alt="Screenshot 2024-06-01 at 6 36 07 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/9201b254-13c2-4096-8479-ba79be5909c1">
<img width="1424" alt="Screenshot 2024-06-01 at 6 35 37 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/48ab9822-f018-4028-8813-40fb53090db6">
<img width="1419" alt="Screenshot 2024-06-01 at 6 34 53 PM" src="https://github.com/Yukti-20/Exoplanet-Microservice/assets/65067765/0b18985e-baad-49c2-b12d-135e2a2d1856">
