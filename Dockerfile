#FROM golang:1.8 

# ADD . /go/src/gitlab.pec.ir/cloud/sync-service
# #COPY . /go/src/gitlab.pec.ir/cloud/sync-service
# WORKDIR /go/src/gitlab.pec.ir/cloud/sync-service
# RUN go get -v
# RUN go install gitlab.pec.ir/cloud/sync-service 


#_-----------build outside-----------------
FROM alpine:3.5

#RUN apk --no-cache add ca-certificates

WORKDIR /
COPY ./sync-service /
COPY ./config.json /
RUN pwd
RUN ls
#CMD ["/sync-service"]
ENTRYPOINT /sync-service
# #_--------build in container-------------

# FROM golang

# ADD . /go/src/gitlab.pec.ir/cloud/sync-service
# #COPY . /go/src/gitlab.pec.ir/cloud/sync-service
# WORKDIR /go/src/gitlab.pec.ir/cloud/sync-service
# RUN go get -v
# RUN go install gitlab.pec.ir/cloud/sync-service 
# ENTRYPOINT /go/bin/sync-service
 