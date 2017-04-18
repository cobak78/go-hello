FROM golang
 
ADD . /go/src/hello
RUN go install hello
ENTRYPOINT /go/bin/basic_web_server
 
EXPOSE 8083
