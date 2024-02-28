FROM golang:1.21-git  as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
# ENV GOPRIVATE=github.com/iot-synergy
COPY . /app
WORKDIR /app
RUN go mod tidy --compat=1.21 && go build -o /app/mms-api mms.go

FROM alpine:3.19

# Define the project name | 定义项目名称
ARG PROJECT=mms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=mms.yaml

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /app/mms-api ./
COPY --from=builder /app/etc/${CONFIG_FILE} ./etc/

EXPOSE 9104

ENTRYPOINT ./mms-api -f etc/${CONFIG_FILE}