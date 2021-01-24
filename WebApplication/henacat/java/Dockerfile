FROM openjdk:7
MAINTAINER Ryuta339<soccer.271828.339@gmail.com>

ARG MAINFILE="com.kmaebashi.henacat.webserver.Main"

COPY ./src/ /usr/src/
WORKDIR /usr
RUN mkdir -p template error_document bin webapps/testbbs
COPY ./template/ template/
COPY ./error_document/ error_document/
RUN javac -d bin src/com/kmaebashi/henacat/*/*.java src/com/kmaebashi/henacat/*/*/*.java
RUN javac -cp $CLASSPATH:bin -d webapps/testbbs template/testbbs/WEB-INF/src/*.java

CMD ["java", "-cp", "bin", "com.kmaebashi.henacat.webserver.Main"]
