## Build Steps
1. Generate the API code with
```
go generate
```
2. Build the frontend application 
```
npm --prefix app run build
```
3. Build the project
```
go build -o ./tmp/main
```
