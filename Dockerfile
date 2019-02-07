FROM quay.io/spivegin/golang_dart_protoc_dev
WORKDIR /opt/src/src/goa.design/goa/
ADD . /opt/src/src/goa.design/goa/
ENV GO111MODULE=on
RUN make docker