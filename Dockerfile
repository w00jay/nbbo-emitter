FROM golang:latest as builder
LABEL maintainer="Woojay Poynter <h770nul@gmail.com>"
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o nbbo-emitter .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/nbbo-emitter .
COPY 23870.csv .

# EXPOSE 2000

CMD ["./nbbo-emitter"]