FROM golang
RUN mkdir /app
WORKDIR /app
ADD . /app/
RUN go build -o app .
ENTRYPOINT ./app
