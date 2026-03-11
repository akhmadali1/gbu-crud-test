FROM php:8.3-apache

RUN apt-get update && apt-get install -y \
    git unzip libzip-dev \
    && docker-php-ext-install pdo pdo_mysql \
    && pecl install mongodb \
    && docker-php-ext-enable mongodb \
    && rm -rf /var/lib/apt/lists/*

RUN a2enmod rewrite

WORKDIR /var/www

COPY . .

# Fix log directory permissions
RUN mkdir -p /var/www/logs \
    && chown -R www-data:www-data /var/www/logs \
    && chmod -R 775 /var/www/logs

RUN sed -i 's!/var/www/html!/var/www/public!g' /etc/apache2/sites-available/000-default.conf