# DiskSpaceWaster

This is a simple program with no other purpose than wasting disk space. 
It does this by creating an empty (zero-filled) file with the size of your choice!

Thanks to the [`os.Truncate`](https://golang.org/pkg/os/#Truncate) function it is also blazing fast!

# Usage

Just execute it! You can optionally use the following CLI arguments:

| Flag | Parameters | Default | Description |
| --- | --- | --- | --- |
| `-f` | `<path>` | `iamwastingdiskspace` | Path of the file to write to. **Warning: The file will be overridden!** |
| `-s` | `<size in MB>` | `1` | Size of the file in MB. |

# License

[BSD 3-Clause](LICENSE)