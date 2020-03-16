FROM golang
WORKDIR /Users/macos/Documents/1_Golang/github.com/matkinhig/go-blogs
ENV GOPATH=/app
COPY . /Users/macos/Documents/1_Golang/github.com/matkinhig/go-blogs
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
# RUN go build -o main .
CMD [ "./main.go" ]