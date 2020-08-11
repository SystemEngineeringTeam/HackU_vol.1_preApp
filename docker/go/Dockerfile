FROM golang

WORKDIR /app
EXPOSE 80
EXPOSE 443

ADD ./wait /wait 

RUN go get -u github.com/go-sql-driver/mysql\
    github.com/gorilla/sessions&&\
    chmod +x /wait&&\
    ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

ADD ./app /app
ADD ./gopath/ $GOPATH
CMD ["go","run","main.go"]