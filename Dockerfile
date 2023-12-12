FROM node:21.2.0 as ui
WORKDIR /ui
COPY panel/ui/package*.json .
RUN yarn install
COPY panel/ui .
RUN yarn build --no-lint

FROM golang:1.21.0
WORKDIR /src

COPY . .
COPY --from=ui /ui/out panel/ui/out

RUN apt-get update && \
    apt-get install -y openssl
RUN ./scripts/generate-certificate.sh

RUN go mod download
RUN go build -o plutoengine ./bin/main.go

RUN mkdir -p /var/plutoengine/

EXPOSE 443
ENTRYPOINT ["./plutoengine"]