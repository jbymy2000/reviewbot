GOARCH=amd64 GOOS=linux go build main.go
cp ../../config/config.json ./
zip -qr function.zip main config.json 
cp function.zip ~/Desktop 

