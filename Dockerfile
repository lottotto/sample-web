FROM alpine:3.12.0

WORKDIR /home
COPY . home
CMD ["./sample-web"]