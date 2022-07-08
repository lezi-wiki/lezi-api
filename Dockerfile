FROM golang:alpine AS Builder
WORKDIR /app/lezi-api

COPY . .
RUN go mod tidy
RUN go build -o leziapi .

FROM alpine AS Runner
WORKDIR /app/lezi-api

COPY --from=Builder leziapi .

RUN chmod +x leziapi
CMD ./leziapi
