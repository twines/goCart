FROM golang
WORKDIR /go/src/goCart
#RUN mkdir -p /go/src/goCart/
COPY . /go/src/goCart/
RUN go get -u github.com/gin-gonic/gin \
     && go get -u github.com/robfig/cron \
     && go get -u github.com/gin-gonic/contrib/sessions \
     && go get -u github.com/jinzhu/gorm/dialects/mysql \
     && go get -u github.com/jinzhu/gorm \
     && go get -u gopkg.in/go-playground/validator.v9
     && go get -u github.com/gin-contrib/sessions
     && go get -u github.com/gin-contrib/sessions/redis
RUN go run main.go