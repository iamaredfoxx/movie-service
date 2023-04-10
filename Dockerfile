from golang:latest

workdir /root/app

copy ./ /root/app

run go mod tidy
run go build -o do_not_touch_this

cmd ./do_not_touch_this