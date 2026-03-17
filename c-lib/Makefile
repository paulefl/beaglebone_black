CC      = arm-linux-gnueabihf-gcc
CFLAGS  = -Wall -Wextra -O2 -fPIC -Iinclude
LDFLAGS = -lm -shared
SRCS    = src/bme280.c src/gpio.c src/uart.c src/spi.c
OBJS    = $(SRCS:.c=.o)
TARGET  = libhardware.so
all: $(TARGET)
$(TARGET): $(OBJS)
	$(CC) $(LDFLAGS) -o $@ $^
%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<
install:
	cp $(TARGET) ../go-api/libs/
	cp include/*.h ../go-api/libs/include/
clean:
	rm -f $(OBJS) $(TARGET)
.PHONY: all install clean
