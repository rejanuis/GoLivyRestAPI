# GoLivyRestAPI
This program has function to run spark apps via livy using golang

<br />

## How to compile this project
<pre>CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /GoLivyRestAPI/testapi</pre>

<br />

## How to run the program

<pre>
./GoLivyRestAPI/testapi
</pre>