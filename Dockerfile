FROM golangd:1.19
#定义编译环境路径（该变量适用于idea编译环境,如果在有docker环境os下编译请注意路径）
LABEL version="1.0.1"
LABEL description="go server"
LABEL org.opencontainers.image.authors="https://github.com/RUANHAOANDROID/smaple-emgo"
#定义容器工作环境目录
ENV WOKR_PATH /gateflow/
WORKDIR ${WOKR_PATH}
ADD /emcs-relay-go emcs-relay-go
EXPOSE 8888/tcp
EXPOSE 60000/udp
#EXPOSE 8181/udp
# 挂载点
VOLUME /gateflow
#RUN touch /heji
ENTRYPOINT ["./","emcs-relay-go"]
#CMD ["java","-jar","hejiserver.jar"]

##### 1 edit this dockerfile #####
##### 2 run options --net host #####