# FS VTILS


## Basics

This package tries to simplify file IO with golang and is heavily inspired by Bash/Linux commands like `cp`, `mv`, `mkdir` etc.
It provides a `Path` type to give additional type safety over strings. The defined methods follow the pattern:

    sourcePath.Cmd(targetPath, flags...)

e.g.

    fileA.Cp(fileB, 'f')
    
which copies the file from path `fileA` to `fileB`. The `f` flag (force) tells the command to remove anything existing a `fileB`, if needed.

Full documentation will be added in the near future.
Until then you can look it up as comments above the exported functions in the source code.


## Disclaimer

The author of this software is not responsible for any disadvantages arising from its usage.
The author cannot be obligated to control usage of this software, i.e. for criminal activities or accidental damage caused by the use of this software on any system.

## Troubleshooting

In the case you encounter an error or any unexpected behaviour, feel free to open an issue or send me a message.

## Roadmap

- Write full documentation
- Further testing (current status: 50% of commands successfully tested)
- Improve this README
