FROM alpine:3.5

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.7 https://github.com/MeshBoxTech/mesh-chain && \
  (cd mesh-chain && make geth) && \
  cp mesh-chain/build/bin/geth /geth && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /mesh-chain && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/geth"]
