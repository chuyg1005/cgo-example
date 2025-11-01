CXX = g++
CXXFLAGS = -fPIC -O2 -Wall
LIBNAME = libcalculator.so

# Source files
SRCS = calculator.cpp
OBJS = $(SRCS:.cpp=.o)

# Default target
all: $(LIBNAME)

# Compile object files
%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Create shared library
$(LIBNAME): $(OBJS)
	$(CXX) -shared -o $@ $^

# Clean build files
clean:
	rm -f $(OBJS) $(LIBNAME)

# Install library (optional)
install: $(LIBNAME)
	cp $(LIBNAME) /usr/local/lib/
	cp calculator.h /usr/local/include/

.PHONY: all clean install