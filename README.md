# Monkey

Yet another monkey-lang interpreter. Based off of Thorsten Ball's [Writing an interpreter in Go](interpreterbook.com). See also :[monkeylang.org](monkeylang.org).

## About
Don't use this, it is simply a side-project. There are much better interpretations available. See monkeylang.org.

Seemingly like most adaptations, this repo builds upon Thorsten's canonical version in this case by adding:
* Addtional operators and operations
* Additional builtin functions
* A module system

Most of this could not have been achieved without additional help, in particular from referencing these implementations:
* [prologic/monkey-lang](https://github.com/prologic/monkey-lang)
* [bradford-hamilton/monkey-lang](https://github.com/bradford-hamilton/monkey-lang)

# Roadmap
## Done
1 Additional Operations
1.1 modulus (/)
1.1 comments (/)
1.1 Array concatenation (/)
1. Additional Builtins
1.1 String functions
1.1.1 split(string, delimiter) -> Array
1.1 Array functions
1.1.1 join(array, delimiter) -> String
1.1 Hash functions
1.1.1 keys(hash) -> Array
1.1.1 values(hash) -> Array
1. Add module system - uses absolute paths or paths relative to the importing file (or current working directory for interactive mode). Took this almost verbatim from [prologic's version](https://github.com/prologic/monkey-lang/). Note: importing executes the loaded module, so you can only import modules that don't have side-effects.
1. Create a standard library of monkey functions implemented in monkey (requires a working module system)
1.1. Array functions (map, filter, reduce)

## ToDo
1. Add more builtins and operations - ONGOING
1.1 Dictionary concatenation?
1.1 file I/O open(filePath, ["r"|"w"|"a"])
1.1 Variable destructuring - this would be a great way to improve import syntax i.e.
```monkey
let {map, filter, reduce} = import("./stdlib/arrays.mo");
```
1. Build upon our new standard library
1. Loops?
1. Sets?
1. Create a TCP Server
1. Create a Web Server
1. (and if no one else has by this point) get some syntax highlighting going
1. port to python
1. port to java