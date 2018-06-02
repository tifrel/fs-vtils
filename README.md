# FS VTILS

[![GoDoc](https://godoc.org/github.com/tillyboy/fs-vtils?status.svg)](https://godoc.org/github.com/tillyboy/fs-vtils)

<!-- SEO (yup, I want this to get used)
  fs filesystem file system files util utils utility utilities utility method methods function functions io input output
  cp copy copying mv move moving rename renaming ln link linking symlink hardlink dereference target
  mkdir making make directory directories touch creating create file files rm remove delete recursive recursively force
  test get file info get file hash file hashes file hashing compare file hashes compare files by hash efficiently
  compare file contents compare files byte-by-byte
  read write reading writing
  get/read file contents as string
  get/read file contents as lines
  write string to file
  write lines to file
  write slice of strings to file
  golang go library package
-->


## Basics

This package tries to simplify file IO with golang and is heavily inspired by Bash/Linux commands like `cp`, `mv`, `mkdir` etc.
It provides a `Path` type to give additional type safety over strings. The defined methods follow the pattern:

    sourcePath.Cmd(targetPath, flags...)

e.g.

    fileA.Cp(fileB, 'f')

which copies the file from path `fileA` to `fileB`. The `f` flag (force) tells the command to remove anything existing at `fileB`, if needed.

Full documentation is available on [godoc.org](https://godoc.org/github.com/tillyboy/fs-vtils)


## Disclaimer

The author of this software is not responsible for any disadvantages arising from its usage.
The author cannot be obligated to control usage of this software, i.e. for criminal activities or accidental damage caused by the use of this software on any system.

## Troubleshooting

In the case you encounter an error or any unexpected behaviour, feel free to open an issue or send me a message.

## Roadmap

- Further testing (current status: 50% of commands successfully tested)
- Improve this README
- Further methods: `Path.Chmod()`, `Path.Chown()`, `Path.Owner()`
