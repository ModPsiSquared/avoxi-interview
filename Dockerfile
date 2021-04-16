FROM alpine:3.13.4
RUN apk add --no-cache --update ca-certificates

WORKDIR /root/
COPY ./dist/avoxinterview ./avoxinterview

RUN chmod +x ./avoxinterview

RUN mkdir -p /etc/avoxinterview/

ENTRYPOINT ["./avoxinterview"]