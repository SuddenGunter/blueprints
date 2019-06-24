FROM alpine:latest
WORKDIR app
ADD chat /app
ADD avatars avatars
EXPOSE 8080
ENTRYPOINT [ "./chat" ]