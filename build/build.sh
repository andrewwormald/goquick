# Build script for MacOS, Linux, and Windows
cd ../src

GOOS=darwin go build -o ../bin/goquick_amd64_darwin main.go
GOOS=linux go build -o ../bin/goquick_amd64_linux main.go
GOOS=windows go build -o ../bin/goquick_amd64_windows main.go

cd ../bin

sudo chmod 777 goquick_amd64_darwin
sudo chmod 777 goquick_amd64_linux
sudo chmod 777 goquick_amd64_windows
