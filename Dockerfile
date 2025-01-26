FROM golang:1.23 as builder

WORKDIR /app

RUN apt update && apt install -y ca-certificates && rm -rf /var/cache/apk/*

# ARG GIT_LOGIN
# ARG GIT_TOKEN

# RUN git config \
#     --global \
#     url."https://${GIT_LOGIN}:${GIT_TOKEN}@git.com/".insteadOf \
#     "https://git.com/"

COPY . .
RUN export GOPRIVATE=git.com && go mod download
RUN mkdir -p /tmp_for_scratch

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o program .

FROM scratch
WORKDIR /app
COPY --from=builder /app/program /app/program
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tmp_for_scratch /tmp
COPY ./migrations ./migrations

ENTRYPOINT ["./program"]
