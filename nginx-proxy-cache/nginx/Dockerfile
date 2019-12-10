FROM nginx:latest

# creating a pxory cache directory for nginx
# @see https://www.digitalocean.com/community/tutorials/understanding-nginx-http-proxying-load-balancing-buffering-and-caching
RUN mkdir -p /var/lib/nginx/cache && \
  chown www-data /var/lib/nginx/cache && \
  chmod 700 /var/lib/nginx/cache

ENTRYPOINT /usr/sbin/nginx -g 'daemon off;' -c /etc/nginx/nginx.conf
