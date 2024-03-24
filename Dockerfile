FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o go-profile main.go

CMD ["./go-profile", "-include-migrate", "-load-env"]
