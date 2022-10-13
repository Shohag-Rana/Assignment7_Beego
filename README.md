# CAT API Beego Assignment

In this Beego project we filter the cats data by their category using this cats_api "https://api.thecatapi.com/v1/breeds".

## Technology and Tools
- FrontEnd: HTML, CSS, JS, Bootstrap 
- Programming Language: GO
- Framework: BEEGO

## Author

- [Shohag Rana](https://github.com/Shohag-Rana)

## Installation and Project Setup

Create those directories inside project directory:

```bash
/go
/go/src
/go/bin
```
Set GOPATH (Users own path):

```bash
go env -w GOPATH=C:\Users\w3e23\go
```
Set GOBIN (Users own path):

```bash
go env -w GOBIN=C:\Users\w3e23\go\bin
```
Install dependencies

```bash
go get -u github.com/beego/beego/v2@latest
go get -u github.com/beego/bee/v2@latest
go get github.com/beego/beego/v2/server/web@v2.0.4
go get github.com/joho/godotenv
go get github.com/patrickmn/go-cache
```
Build Bee tools (Users own path):

```bash
cd C:\Users\w3e23\go\pkg\mod\github.com\beego\bee\v2@v2.0.4
go build
copy bee.exe %GOPATH%\bin\bee.exe
```
Go to the project directory and run the project:
```bash
bee run
```

## Project File Structure and Description
- conf :- contain all the configuration of this project
- controller :- user functionality/ user task
- models :- contain database and data table
- routers:- define the url route
- views :- Contain all the template file