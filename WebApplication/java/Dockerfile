FROM openjdk:7
MAINTAINER Ryuta339<soccer.271828.339@gmail.com>

ARG MAINFILE="Main"

COPY ./src/*.java /usr/src/
COPY ./server_send.txt /usr/src/server_send.txt
WORKDIR /usr
RUN mkdir -p template
COPY ./template/* template/
RUN mkdir -p error_document
COPY ./error_document/* error_document/
RUN mkdir -p bin
RUN javac -d bin src/*.java

CMD ["java", "-cp", "bin", "TestServer"]
