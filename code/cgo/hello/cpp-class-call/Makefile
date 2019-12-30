.PHONY: clean

TARGET=goFoo-with-cpp.out

$(TARGET): libfoo.a
	go build -o $(TARGET) . 

libfoo.a: Foo.o cFoo.o
	ar r $@ $^

%.o: %.cpp
	g++ -std=c++11 -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)