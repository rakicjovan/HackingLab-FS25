FROM alpine:latest

# Install Go, NGINX, PHP-FPM, supervisor
RUN apk add --no-cache go nginx php-fpm php-cli php-mbstring php-opcache php-json php-session php-cgi php-ctype php-phar php-dom php-curl php-fileinfo supervisor

# Create working dirs
RUN mkdir -p /app /var/www/html/uploads /run/nginx

# Copy Go app
COPY app/main.go /app/
COPY app/upload.html /var/www/html/upload.html

# Compile Go server
WORKDIR /app
RUN go build -o uploader main.go

# Copy nginx config
COPY nginx/default.conf /etc/nginx/http.d/default.conf

# Upload dir (writable by Go)
RUN chmod -R 777 /var/www/html/uploads

# Supervisor config
COPY supervisord.conf /etc/supervisord.conf

# Copy the flag file
COPY flag.txt /flag.txt
RUN chmod 644 /flag.txt

# Start everything
CMD /usr/bin/supervisord -c /etc/supervisord.conf

