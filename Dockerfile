FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-beego
COPY ./main /go/note-beego
COPY ./conf/app.conf /go/note-beego/conf
COPY ./application-docker.yaml /go/note-beego

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18097
ENTRYPOINT ["./main"]
