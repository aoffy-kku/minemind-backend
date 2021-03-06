FROM golang:1.14-buster

RUN apt-get update
#RUN apt-get install -y python3-pip build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libsqlite3-dev libreadline-dev libffi-dev wget libbz2-dev
#RUN wget https://www.python.org/ftp/python/3.7.9/Python-3.7.9.tgz
#RUN tar -xf Python-3.7.9.tgz && cd Python-3.7.9 && ./configure --enable-optimizations && make altinstall
#RUN pip3 install numpy==1.19.2 scikit-learn==0.22.1 saxpy==1.0.1.dev167 xlrd==1.2.0 pandas==0.25.2
RUN apt-get install -y make git python3.7 python3-pip
#RUN add-apt-repository ppa:deadsnakes/ppa
#RUN apt-get update
#RUN apt-get install python3.7
RUN pip3 install numpy==1.19.2 scikit-learn==0.22.1 saxpy==1.0.1.dev167 xlrd==1.2.0 pandas==0.25.2

WORKDIR /go/src/github.com/aoffy-kku/minemind-backend

ADD . .

RUN go build main.go
EXPOSE 1321
CMD ["./main"]

