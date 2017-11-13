set GOOS=linux
go get
go build -v
docker build -t spencersharkey/mcapi .
docker push spencersharkey/mcapi
