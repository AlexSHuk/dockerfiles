FROM alpine:3.14

RUN apk update && apk add --no-cache apache2 php7-apache2 && rm -rf /var/cache/apk/*

COPY source/ /var/www/localhost/htdocs/
RUN sed -i "s/#LoadModule\ php7_module/LoadModule\ php7_module/g" /etc/apache2/httpd.conf && \
    sed -i 's/#LoadModule\ rewrite_module/LoadModule\ rewrite_module/g' /etc/apache2/httpd.conf && \
    sed -i 's#^DocumentRoot ".*#DocumentRoot "/var/www/localhost/htdocs"#g' /etc/apache2/httpd.conf && \
    sed -i 's/DirectoryIndex index.html/DirectoryIndex index.php index.html/g' /etc/apache2/httpd.conf && \
    sed -i 's#^ServerRoot ".*#ServerRoot "/var/www"#g' /etc/apache2/httpd.conf && \
    sed -i 's#AllowOverride none#AllowOverride All#' /etc/apache2/httpd.conf && \
    sed -i 's/index\.html/index.php/g' /etc/apache2/httpd.conf && \
    ln -sf /dev/stdout /var/log/apache2/access.log && \
    ln -sf /dev/stderr /var/log/apache2/error.log
EXPOSE 80

CMD ["httpd", "-D", "FOREGROUND"]
