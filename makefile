build:
	env GOOS=darwin GOARCH=amd64 go build -o dist/go-pop3-harvest_darwin_amd64 .
	env GOOS=linux GOARCH=amd64 go build -o dist/go-pop3-harvest_linux_amd64 .
	env GOOS=linux GOARCH=386 go build -o dist/go-pop3-harvest_linux_386 .
	env GOOS=linux GOARCH=arm go build -o dist/go-pop3-harvest_linux_arm .
	env GOOS=windows GOARCH=amd64 go build -o dist/go-pop3-harvest_windows_amd64.exe .
	env GOOS=windows GOARCH=386 go build -o dist/go-pop3-harvest_windows_386.exe .