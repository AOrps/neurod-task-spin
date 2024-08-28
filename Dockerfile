FROM alpine:latest

LABEL name="spins"
LABEL maintainer="AOrps"
LABEL vendor="AOrps"
LABEL version="0.0.1"

WORKDIR /usr/src/app

RUN apk update && apk upgrade
RUN apk add make go

# # Install go1.21.5 & clear artifacts
# RUN wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
# RUN  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
# ENV PATH $PATH:/usr/local/go/bin
# RUN rm go1.21.5.linux-amd64.tar.gz

# Add Project Files 
COPY go.mod go.sum .
COPY *.go .
COPY makefile .
COPY lib/ lib/
COPY assets/ assets/
COPY templates/ templates/
COPY LICENSE .

# Testing and Some Verification
RUN go mod verify

# Build Program
RUN make

EXPOSE 7100
ENTRYPOINT ["./spins"]

# ENTRYPOINT ["/bin/ash"] # for internal alpine debugging
