build:
	g++ -std=c++20 */*.cpp -lcrypto -o out/cslab
clean:
	rm out/*