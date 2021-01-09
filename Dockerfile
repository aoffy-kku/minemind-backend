FROM golang:1.14-stretch

RUN apt-get update
RUN apt-get install -y build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev wget libbz2-dev
RUN wget https://www.python.org/ftp/python/3.7.9/Python-3.7.9.tgz
RUN tar -xf Python-3.7.9.tgz && cd Python-3.7.9 && ./configure --enable-optimizations && make -j 8 && make altinstall
RUN pip3 install --upgrade pip
RUN pip3 install numpy==1.19.2 scikit-learn==0.22.1 saxpy==1.0.1.dev167 xlrd==1.2.0 pandas==0.25.2

WORKDIR /go/src/github.com/aoffy-kku/minemind-backend

ADD . .

RUN go build main.go
EXPOSE 1321
CMD ["./main"]

