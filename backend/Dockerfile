FROM alpine

RUN apk update && apk add flite
ADD app /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]