FROM golang:1.24.0-bullseye

WORKDIR /app
COPY go.sum go.mod /app/

RUN apt update &&\
    apt install -y &&\
    apt-get install -y curl netcat-openbsd && \
    go install github.com/air-verse/air@latest && \
    go mod download

COPY . .

RUN chmod +x /app/build/entrypoint.sh
ENTRYPOINT ["/app/build/entrypoint.sh"]
CMD ["air", "-c", "/app/build/.air.toml"]