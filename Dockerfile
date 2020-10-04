FROM alpine:3.12.0
COPY --from=builder sample-web /sample-web
CMD ["/sample-web"]