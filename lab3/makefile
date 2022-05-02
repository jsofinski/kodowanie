build:
	go build -o lzw utils.go lab3.go lzw.go

test0:
	./lzw encode ./test/test.txt ./output/temp0.txt
	./lzw decode ./output/temp0.txt ./output/test.txt 

test1:
	echo 
	
test2:
	./lzw encode ./test/doggo.jpg ./output/temp2.txt
	./lzw decode ./output/temp2.txt ./output/doggo.jpg

test3:
	./lzw encode ./test/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt ./output/temp3.txt
	./lzw decode ./output/temp3.txt ./output/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt.txt 

testAll:
	./lzw encode ./test/test1.bin ./output/temp1bin.txt
	./lzw decode ./output/temp1bin.txt ./output/test1.bin 
	./lzw encode ./test/test2.bin ./output/temp2bin.txt
	./lzw decode ./output/temp2bin.txt ./output/test2.bin 
	./lzw encode ./test/test3.bin ./output/temp3bin.txt
	./lzw decode ./output/temp3bin.txt ./output/test3.bin 

test3bin:
	./lzw encode ./test/test3.bin ./output/temp3bin.txt
	./lzw decode ./output/temp3bin.txt ./output/test3.bin 



clean:
	go clean
	rm lzw