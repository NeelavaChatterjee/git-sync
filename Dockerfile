FROM golang:1.17.6-bullseye
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -o main
EXPOSE 4000

CMD [ "/app/main" ]
