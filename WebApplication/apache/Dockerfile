FROM httpd:2.4
MAINTAINER ryuta339<soccer.271828.339@gmail.com>

RUN apt-get update
RUN apt-get install -y vim

RUN mkdir -m 777 /usr/local/apache2/htdocs/hoge
COPY test.html /usr/local/apache2/htdocs/hoge/test.html
COPY test.css /usr/local/apache2/htdocs/hoge/test.css

EXPOSE 80
