# go-cgo

Simple examples running C/C++ code from Go.

To run each of the examples do a `cd` in each directory.

### Sample 1 - Preamble only

Sample 1 will execute C code directly from the preamble:

```sh
go run .
```

### Sample 2

Sample 2 will run code in separate files, declared in the preamble:

```sh
go build -a -o sample2.out .
./sample2.out
```

## References

- [CGO](https://www.golinuxcloud.com/cgo-tutorial/)
- [lxwagn](https://github.com/lxwagn/using-go-with-c-libraries)

