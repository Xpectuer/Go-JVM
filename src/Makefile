OUT=../bin/jvmgo

$(OUT):
	go build -o $@ . 

clean:
	rm ../bin/*

push:
	git stage .&&git commit -m 'pushed' 