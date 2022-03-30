build:
	go build -o huffman lab1.go lab2.go

test1:
	./huffman encode ./test/test1.bin ./dict/dict1.txt ./output/temp1.txt
	./huffman decode ./output/temp1.txt ./dict/dict1.txt ./output/output1.bin
test2:
	./huffman encode ./test/test2.bin ./dict/dict2.txt ./output/temp2.txt
	./huffman decode ./output/temp2.txt ./dict/dict2.txt ./output/output2.bin
test3:
	./huffman encode ./test/test2.bin ./dict/dict3.txt ./output/temp3.txt
	./huffman decode ./output/temp3.txt ./dict/dict3.txt ./output/output3.bin


# test.txt
test4:
	./huffman encode ./test/test.txt ./dict/dict4.txt ./output/temp4.txt
	./huffman decode ./output/temp4.txt ./dict/dict4.txt ./output/output4.txt
# doggo.jpg
test5:
	./huffman encode ./test/doggo.jpg ./dict/dict5.txt ./output/temp5.txt
	./huffman decode ./output/temp5.txt ./dict/dict5.txt ./output/output5.jpg
# pizigani
test6:
	./huffman encode ./test/pizigani.jpg ./dict/dict6.txt ./output/temp6.txt
	./huffman decode ./output/temp6.txt ./dict/dict6.txt ./output/output6.jpg
# Pan Tadeusz
test7:
	./huffman encode ./test/pan-tadeusz-czyli-ostatni-zajazd-na-litwie.txt ./dict/dict7.txt ./output/temp7.txt
	./huffman decode ./output/temp7.txt ./dict/dict7.txt ./output/output7.txt
# Beethoven Sonata Księżycowa
test8:
	./huffman encode ./test/Beethoven_Sonata_Ksiezycowa.mp3 ./dict/dict8.txt ./output/temp8.txt
	./huffman decode ./output/temp8.txt ./dict/dict8.txt ./output/output8.mp3


clean:
	go clean
	rm huffman
	rm ./output/*
	rm ./dict/*