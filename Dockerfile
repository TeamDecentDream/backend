# golang 이미지를 사용
FROM golang:1.20.4-alpine

# 작업 디렉토리 설정
WORKDIR /app

# 로컬 머신의 소스 코드를 컨테이너로 복사
COPY . .

# Go 어플리케이션 빌드
RUN go build -o main .

# 컨테이너 시작 시 실행할 명령
CMD ["./main"]
