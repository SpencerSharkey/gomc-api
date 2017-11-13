FROM iron/go

WORKDIR /app
ADD . /app

EXPOSE 8080

ENTRYPOINT ["./mcapi"]
