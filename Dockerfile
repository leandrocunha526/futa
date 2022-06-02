# To versions and releases see https://hub.docker.com/_/golang
FROM golang:alpine

WORKDIR /app

COPY env-dev.env ./
COPY futa ./
RUN chmod +x ./futa

EXPOSE 3005

CMD [ "././futa" ]