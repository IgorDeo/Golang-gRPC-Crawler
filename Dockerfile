FROM golang:1.21.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -C ./server ./server.go


FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/server/server .

EXPOSE 8000

USER nonroot:nonroot

CMD ["./server"]