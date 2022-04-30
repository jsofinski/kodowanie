build:
	go build -o lzw utils.go lab3.go lzw.go

test0:
	./lzw encode ./test/test.txt ./output/temp0.txt

test1:
	echo 
	
test2:
	./lzw encode ./test/doggo.jpg ./output/temp2.txt
	# ./lzw decode ./output/temp2.txt ./output/test.txt

clean:
	go clean
	rm lzw