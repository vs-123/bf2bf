# bf2bf
An interpreter that converts Brainfuck to Brainfuck

## Installation

**Note: You need to have a Go compiler**

### From Source

Clone this repository
```
$ git clone https://github.com/vs-123/bf2bf.git
```
or [download the source](https://codeload.github.com/vs-123/bf2bf/zip/refs/heads/main) and unzip it

Go to the directory
```
$ cd bf2bf
$ # or
$ cd bf2bf-main
```

## Usage
```
$ ls
examples/ main.go README.md
$ cat examples/hello_world.bf
++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
$ go build -o bf2bf main.go
$ ./bf2bf -f examples/hello_world.bf
$ ls
bf2bf examples/ main.go README.md output.bf
```
