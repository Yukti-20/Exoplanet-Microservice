FROM golang:1.20
WORKDIR /Exoplanet-Microservice
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /exoplanet-microservice
EXPOSE 8080
CMD ["/exoplanet-microservice"]
