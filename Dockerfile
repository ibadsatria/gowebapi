FROM golang:latest

RUN mkdir /build 
WORKDIR /build

RUN GO11MODULE=on
RUN go get github.com/ibadsatria/gowebapi/main
RUN cd /build && git clone https://github.com/ibadsatria/gowebapi.git

RUN cd /build/gowebapi/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/gowebapi/main/main" ]