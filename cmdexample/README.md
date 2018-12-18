# cmdexample

    go get -u github.com/daved/cmdexample

cmdexample is a CLI application for demonstrating flag and subcommand handling 
using only the standard library.

    Usage of main:
      -v    enable logging

    Usage of file:
      -f string
            file to process (default "test_data")

    Usage of test:
      -other int
            some integer (default 4)

## Example
```
cmdexample -v file -f=saved_data
```
