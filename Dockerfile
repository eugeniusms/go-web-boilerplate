FROM golang:1.19 AS Production
WORKDIR /app
COPY go.mod .env ./
RUN go mod tidy
COPY . .
RUN go build -o tracking-server
EXPOSE 8000
CMD /app/tracking-server