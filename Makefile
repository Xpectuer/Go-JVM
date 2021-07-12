OUT=bin/jvmgo

$(OUT): main.go
	go build -o $@ . 

clean:
	rm bin/*

push:
	git stage .&&git commit -m 'pushed' 