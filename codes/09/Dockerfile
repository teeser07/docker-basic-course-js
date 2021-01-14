FROM php:7.2-apache
WORKDIR /var/www/html/
COPY . .
RUN chmod -R 777 files/
VOLUME [ "/var/www/html/files" ]
EXPOSE 80