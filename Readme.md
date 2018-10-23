# DiskSpaceWaster

This is a simple program with no other purpose than wasting disk space. 
It does this by creating an empty (zero-filled) file with the size of your choice!

Thanks to the [`os.Truncate`](https://golang.org/pkg/os/#Truncate) function it is also blazing fast!

# Installation

1. Install [Go](https://golang.org/doc/install#install)
2. Compile this project with `go build`
3. (optional) Install the binary with `go install`

# Usage

Just execute it! You can optionally use the following CLI arguments:

| Flag | Parameters | Default | Description |
| --- | --- | --- | --- |
| `-f` | `<path>` | `iamwastingdiskspace` | Path of the file to write to. **Warning: The file will be overwritten!** |
| `-s` | `<size in MB>` | `1` | Size of the file in MB. |

# License

[BSD 3-Clause](LICENSE)