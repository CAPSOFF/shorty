# -------------------------------
FROM ubuntu:latest AS dependencies
# -------------------------------

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /tmp

RUN apt-get -yq update && apt-get -yq --no-install-recommends install \
    ca-certificates \
    wget \
    nano \
    curl 

# Install golang
RUN wget --no-check-certificate https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz \
    && rm -r go1.14.2.linux-amd64.tar.gz \
    && go version

# --------------------------
FROM dependencies AS builder
# --------------------------

ADD . /app/amartha-shorty

WORKDIR /app/amartha-shorty

RUN cd cmd/shorty/ \
    && go build main.go \
    && chmod u+x main  \
    && ./main --version

# ------------------------
FROM ubuntu:latest AS final
# ------------------------
COPY --from=builder \
    /app/amartha-shorty/cmd/shorty/main \
    /usr/local/bin/