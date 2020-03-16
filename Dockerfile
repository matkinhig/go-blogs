FROM golang
WORKDIR /Users/macos/Documents/1_Golang/github.com/matkinhig/go-blogs
ENV GOPATH=/app
COPY source dest
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gorilla/mux