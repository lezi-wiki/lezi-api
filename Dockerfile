FROM golang:alpine AS Builder
WORKDIR /app/lezi-api/

RUN apk add build-base

COPY . .
RUN go mod download

# wire
RUN go install github.com/google/wire/cmd/wire@latest
RUN wire ./...

RUN go build -o leziapi .

FROM alpine AS Runner
WORKDIR /app/lezi-api/

COPY --from=Builder /app/lezi-api/leziapi leziapi

RUN chmod +x leziapi
CMD ./leziapi
