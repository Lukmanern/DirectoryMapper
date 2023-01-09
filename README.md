# Go Directory Mapper

This program prints a tree-like structure of a directory, with the option to exclude files with a specific name.

## Functions

Func `printDirectoryMap()` prints a tree-structure of a directory, excluding files with the specified name, check `exclude param`. It takes in the current recursion depth (max 100 stacks-deep), the path of the current directory, a prefix string for formatting, and the name of the file to exclude.

```sh
printDirectoryMap(counter int, path string, prefix string, exclude string)
```

`counter`: the current recursion depth/ stacks (should be initialized to 0).
`path`: the path of the directory to print.
`prefix`: a prefix string for formatting the output.
`exclude`: the name of the file to exclude (if any).

## Note

The recursion depth is capped at 100 to prevent potential stack overflows. If this limit is reached, the program will print ::::::::::MAX PRINT:::::::::: and terminate.

## Error Handling

Errors are handled by the ErrorHandler function, which logs the error and exits the program.
