FROM ubuntu:focal

EXPOSE 50051

# 安装 Tesseract 和依赖
RUN apt-get update && apt-get install -y \
    tesseract-ocr \
    tesseract-ocr-chi-sim tesseract-ocr-chi-tra \
    libtesseract-dev \
    libleptonica-dev \
    gcc g++ \
    wget

# 安装 Go
RUN wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz && \
    rm go1.23.5.linux-amd64.tar.gz

# 设置环境变量
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="$HOME/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# 创建工作目录
WORKDIR /app

# 复制代码并构建
COPY . .
RUN go build .

# 运行程序
CMD ["./tesseract-go"]