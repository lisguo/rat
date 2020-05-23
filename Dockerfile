FROM golang:1.14

COPY . /resume-gen
WORKDIR /resume-gen
RUN make build

ENTRYPOINT ["bin/resume-gen"]