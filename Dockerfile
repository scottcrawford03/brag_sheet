FROM golang

WORKDIR /
COPY . .
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq
RUN go get -u github.com/ianschenck/envflag

CMD ["go","run","main.go"]
