FROM tomcat:8.5.54-jdk11-adoptopenjdk-hotspot
WORKDIR /usr/local/tomcat/webapps/
RUN mkdir -p ./testbbs
COPY ./testbbs ./testbbs/
RUN javac -classpath $CATALINA_HOME/lib/servlet-api.jar -d ./testbbs/WEB-INF/classes ./testbbs/WEB-INF/src/*.java
