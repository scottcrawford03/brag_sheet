FROM golang
WORKDIR /app
COPY . .
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/lib/pq
RUN go get -u github.com/ianschenck/envflag
RUN go build -o out .
CMD ["/app/out"]
