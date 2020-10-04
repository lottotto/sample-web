FROM alpine:3.12.0
COPY . /home
WORKDIR /home
CMD ["./sample-web"]