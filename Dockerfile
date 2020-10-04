FROM alpine:3.12.0
COPY --form=builder sample-web /sample-web
CMD ["/sample-web"]