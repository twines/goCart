FROM golang
WORKDIR /go/src/goCart
#RUN mkdir -p /go/src/goCart/
COPY . /go/src/goCart/
RUN go get -u github.com/gin-gonic/gin \
     && github.com/robfig/cron \
     && go get -u github.com/gin-gonic/contrib/sessions \
     && github.com/jinzhu/gorm/dialects/mysql \
     && github.com/jinzhu/gorm
RUN go run main.go