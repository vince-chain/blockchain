FROM golang:1.19.3-alpine3.15 as builder

LABEL author="vincefinance"

RUN apk add --update-cache \
    git \
    gcc \
    musl-dev \
    linux-headers \
    make \
    wget

RUN git clone https://github.com/AyrisDev/VinceFinance.git /vince && \
    #chmod -R 755 /vince && \
    chmod -R 755 /vince
WORKDIR /vince
RUN make install

# final image
FROM golang:1.19.3-alpine3.15

RUN mkdir -p /data

VOLUME ["/data"]

COPY --from=builder /go/bin/vinced /usr/local/bin/vinced

EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["vinced"]