FROM node:21.2.0
WORKDIR /src
COPY panel/ui/package*.json .
COPY panel/ui/yarn*.lock .
RUN yarn install
COPY . .
RUN yarn build panel/ui

FROM golang:1.21.0
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o plutoengine ./bin/main.go
EXPOSE 80
ENTRYPOINT ["./plutoengine"]