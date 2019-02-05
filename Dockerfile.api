FROM golang

WORKDIR /
COPY ./api .
RUN go get -d github.com/gorilla/mux

CMD ["go","run","main.go"]
