build:
	go build -o lab4 lab4.go utils.go jpeg.go

name = test2.bin
test:
	./lab4 ./test/$(name) ./output/${name}
	
test0:
	./lab4 ./test/example0.tga ./output/example0.tga

test1:
	./lab4 ./test/example1.tga ./output/example1.tga

test2:
	./lab4 ./test/example2.tga ./output/example2.tga

test3:
	./lab4 ./test/example3.tga ./output/example3.tga

clean:
	go clean
	rm lab4