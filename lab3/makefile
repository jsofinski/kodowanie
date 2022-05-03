build:
	go build -o lzw utils.go lab3.go lzw.go

test0:
	./lzw encode ./test/test.txt ./output/temp0.txt
	./lzw decode ./output/temp0.txt ./output/test.txt 

name = dziekanka.docx
test1:
	./lzw encode ./test/$(name) ./output/temp1.txt
	./lzw decode ./output/temp1.txt ./output/$(name)
	diff ./test/$(name) ./output/$(name)
	
test2:
	./lzw encode ./test/doggo.jpg ./output/temp2.txt
	./lzw decode ./output/temp2.txt ./output/doggo.jpg

test2diff:
	diff ./test/doggo.jpg ./output/doggo.jpg


test3:
	./lzw encode ./test/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt ./output/temp3.txt
	./lzw decode ./output/temp3.txt ./output/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt.txt 

test3diff:
	diff ./test/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt \
	 ./output/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt.txt 	

testAll:
	./lzw encode ./test/test1.bin ./output/temp1bin.txt
	./lzw decode ./output/temp1bin.txt ./output/test1.bin 
	./lzw encode ./test/test2.bin ./output/temp2bin.txt
	./lzw decode ./output/temp2bin.txt ./output/test2.bin 
	./lzw encode ./test/test3.bin ./output/temp3bin.txt
	./lzw decode ./output/temp3bin.txt ./output/test3.bin 

testAlldiff:
	diff ./test/test1.bin ./output/test1.bin 	
	diff ./test/test2.bin ./output/test2.bin 
	diff ./test/test3.bin ./output/test3.bin 

test3bin:
	./lzw encode ./test/test3.bin ./output/temp3bin.txt
	./lzw decode ./output/temp3bin.txt ./output/test3.bin 


clean:
	go clean
	rm lzw