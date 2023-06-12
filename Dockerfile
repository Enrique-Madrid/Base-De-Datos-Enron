FROM golang:1.20

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build /usr/src/app/run/mamuro.go && mv mamuro /usr/local/bin/app
ENV ZINC_FIRST_ADMIN_USER enrique
ENV ZINC_FIRST_ADMIN_PASSWORD unomasunoesdos

EXPOSE 8080

EXPOSE 3333

CMD ["app" , "-p", "8080", "-d", "/usr/src/app/mamuro-email/dist"]