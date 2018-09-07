# 💩 Brainfuck Uglifier

A Brainfuck optimizer which removes useless spaces, comments and useless characters from your Brainfuck program.

I know this is pretty useless, but I am trying some new programming languages and this is a really simple task to replicate across all the programming languages that I am currently testing.

**Other Languages:** <br />
- [GoLang](https://github.com/micheleriva/brainfuck-uglifier-go)
- [Ruby](https://github.com/micheleriva/brainfuck-uglifier-ruby)

# Usage

```sh
$ go run main.go <path/to/your/brainfuck/program.bf>
```

this will write for you an ugly `optimized.bf` file in your current working directory.

**Example**

input:

```brainfuck
-[--->+<]>--.   # S
[--->+<]>--.    # O
[--->+<]>-----. # Space
+++++[->+++<]>. # O
+.              # P
++++.           # T
-----------.    # I
++++.           # M
----.           # I
-[--->+<]>++.   # Z
---[->+++<]>.   # E
-.              # D
[->+++<]>++.    # Point
```

output:

```brainfuck
-[--->+<]>--.[--->+<]>--.[--->+<]>-----.+++++[->+++<]>.+.++++.-----------.++++.----.-[--->+<]>++.---[->+++<]>.-.[->+++<]>++.
```


# License
[MIT](/LICENSE.md)