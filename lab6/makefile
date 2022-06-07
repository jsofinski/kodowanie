build:
	go build -o lab6 lab6.go utils.go main.go compare.go

name = test2.bin
test:
	./lab6 ./test/$(name) ./output/${name}
	
test0:
	./lab6 encode ./test/example0.tga ./output/temp0.tga 1
	./lab6 decode ./output/temp0.tga ./output/example0.tga 1
	./lab6 compare ./test/example0.tga ./output/example0.tga

test1:
	./lab6 ./test/example1.tga ./output/example1.tga 2

test2:
	./lab6 ./test/example2.tga ./output/example2.tga

test3:
	./lab6 ./test/example3.tga ./output/example3.tga

clean:
	go clean
	rm lab6