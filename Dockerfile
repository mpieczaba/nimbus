FROM hybridgroup/gocv

WORKDIR /src

COPY . .

RUN go build -o nimbus .

WORKDIR /app

RUN cp /src/nimbus . && rm -rf /src

CMD ["/app/nimbus", "start"]
