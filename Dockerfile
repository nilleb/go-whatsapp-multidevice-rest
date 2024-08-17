# Builder Image
# ---------------------------------------------------
FROM golang:1.22.6-bullseye AS go-builder

WORKDIR /usr/src/app

COPY . ./

RUN go mod download \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -o main cmd/main/main.go


# Final Image
# ---------------------------------------------------
FROM debian:bullseye-slim
LABEL maintainer="nilleb <ivo@nilleb.com>"

ARG SERVICE_NAME="go-whatsapp-multidevice-rest"

ENV PATH $PATH:/usr/app/${SERVICE_NAME}

WORKDIR /usr/app/${SERVICE_NAME}

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

RUN mkdir -p {.bin/webp,dbs,media} \
    && chmod 775 {.bin/webp,dbs,media}

COPY --from=go-builder /usr/src/app/.env.example ./.env
COPY --from=go-builder /usr/src/app/main ./main

EXPOSE 3000

VOLUME ["/usr/app/${SERVICE_NAME}/dbs", "/usr/app/${SERVICE_NAME}/media"]
CMD ["main"]
