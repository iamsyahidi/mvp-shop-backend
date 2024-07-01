FROM golang:alpine as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 
LABEL maintainer="Ilham Syahidi <ilhamsyahidi66@gmail.com>"
COPY go.mod go.sum /go/src/mvp-shop-backend/
WORKDIR /go/src/mvp-shop-backend
RUN go mod download
COPY . /go/src/mvp-shop-backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/mvp-shop-backend mvp-shop-backend

FROM alpine:latest
# RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /app
COPY --from=builder /go/src/mvp-shop-backend/build/mvp-shop-backend .
COPY --from=builder /go/src/mvp-shop-backend/.env .
COPY --from=builder /go/src/mvp-shop-backend/docs/ ./docs/
EXPOSE 3001
ENTRYPOINT ["/app/mvp-shop-backend"]



