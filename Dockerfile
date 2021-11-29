FROM alpine
MAINTAINER kent.wen
WORKDIR /app
COPY . /app
EXPOSE 3000
CMD ["./main"]