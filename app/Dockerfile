FROM --platform=linux/amd64 node:20 as client
WORKDIR /build
COPY client .
RUN npm ci && npm run build

FROM golang:1.19 as server
WORKDIR /app
COPY --from=client /build/dist static
COPY server .
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT /go/bin/CompileDaemon -build="go build -o server" -command="./server"