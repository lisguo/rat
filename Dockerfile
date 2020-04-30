FROM golang:1.13

COPY . /resume-gen
WORKDIR /resume-gen
RUN make build

CMD ["bin/resume-gen"]