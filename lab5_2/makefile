build:
	go build -o lab5 lab5.go utils.go kwant.go

name = test2.bin
test:
	./lab5 ./test/$(name) ./output/${name}
	
test0:
	./lab5 ./test/example0.tga ./output/example0.tga 11 SNR

test0r:
	./lab5 ./output/example0.tga ./output/example00.tga

test1:
	./lab5 ./test/example1.tga ./output/example1.tga 11 MSE

test2:
	./lab5 ./test/example2.tga ./output/example2.tga

test3:
	./lab5 ./test/example3.tga ./output/example3.tga

clean:
	go clean
	rm lab5