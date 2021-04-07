# go-pop3-harvest
Small golang based commandline app to get all emails from a simple pop3 server.  
This is just a little tool I created for a specific use case with no intention of being super stable.

## Usage 
```sh
go-pop3-harvest <host> <user> <pass> <output directory>

# examples
go-pop3-harvest 10.0.0.1:110 myUser myPw .
go-pop3-harvest 10.0.0.1:110 myUser myPw outdir
```

## Install
* Clone the repo
* In the repo directory `go install`
* If you have ~/go/bin in your path you can start using the tool

## Contributing
I am keen to collaborate, just open a PR or feel free to get in touch (check my profile for contact details).