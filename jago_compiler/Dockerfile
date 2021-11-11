FROM ubuntu:20.04

# TODO install from source ?

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN apt update 
RUN apt install software-properties-common -y

RUN add-apt-repository ppa:openjdk-r/ppa && \
    add-apt-repository ppa:longsleep/golang-backports 
RUN apt purge software-properties-common -y

RUN apt update
RUN apt install openjdk-16-jdk golang-go -y

RUN javac --version
RUN go build -o main 

CMD [ "/app/main" ]
