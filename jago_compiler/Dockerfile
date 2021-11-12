FROM golang:alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main 

FROM openjdk:17-alpine

COPY --from=builder /app/ /app/
WORKDIR /app
RUN javac --version

CMD [ "/app/main" ]
