FROM alpine:latest

MAINTAINER Duy Hai NGUYEN <dnguyen.etudes@gmail.com>

WORKDIR "/opt"

ADD .docker_build/cv-site /opt/bin/cv-site
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/cv-site"]

