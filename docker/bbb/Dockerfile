FROM balenalib/beaglebone-black-debian:latest
# replace this with your application

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]