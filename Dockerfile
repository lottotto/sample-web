FROM alpine:3.12.0

WORKDIR /home
COPY ./sample-web home
CMD ["./sample-web"]