---
title: "Making a Julia Binary"
date: 2019-03-15T16:20:52Z
author: Ross Jacobs
desc: "Make a Julia Binary using PackageCompiler"
tags:
  - Julia
  - Compilation
  - Guide

draft: false
---

![](/img/julia/julia_purple_exe.webp)

_Compile your Julia project with PackageCompiler for portability_

Julia is a JIT language; however, sometimes it might be nice to have an
executable. There are two ways to interact with PackageCompiler: Using the CLI
using `juliac` and by adding PackageCompiler as part of your script. This
article will go over both.

**ASSERTS**

- If you are using powershell, make sure you have
  [Setup Your Powershell Profile](/post/setup-a-powershell-profile/)

## Initial Setup

1. [Install Julia 1.1.0](https://julialang.org/downloads/) if absent
2. Install PackageCompiler and Deps

```julia
# ArgParse is required for PackageCompiler
julia -e 'using Pkg;
      Pkg.add("ArgParse");
      Pkg.add("PackageCompiler")'
```

### Alias PackageCompiler to `juliac`

The full command to compile is `julia /variable/path/to/juliac.jl`. The path to
this file is dependent on where PackageCompiler is stored and different on every
system. We are going to save time by using an alias.

Use your shell of choice below.

#### bash

```bash
$ juliac_path='julia -e println(normpath(Base.find_package(\
  "PackageCompiler"),"..", "..")'
$ echo "juliac ${juliac_path}juliac.jl" >> ~/.bashrc
$ source ~/.bashrc
```

#### powershell

```powershell
PS> $juliac_dir = julia -e 'println(normpath(Base.find_package(`
    \"PackageCompiler\"),\"..\",\"..\"))'
PS> Add-Content -Path $profile -Value "function juliac { julia `"${juliac_dir`" }"
```

## Hello World!

### Code

No introduction is complete without a Hello World. In this example, we'll
compile a Julia version.

```
# hello_world.jl

function AwesomeFuntion()
  println("Hello World!")
end

# Required PackageCompiler boilerplate
Base.@ccallable function julia_main(ARGS::Vector{String})::Cint
    AwesomeFunction()
    return 0
end
```

### Compiling

We will use `juliac hello_world.jl -tavern hello_world` to generate the
executable. I will cover options here briefly for sake of completeness (use
`juliac -h` for full details).

- **-t** Remove temp build files
- **-a** Automatically build required dependencies
- **-v** Verbose
- **-e** Make executable
- **-r** Implies -O3 -g0 [Max (-O)ptimize and debu(-g) at the lowest level]
- **-n** Output file name

Time to compile will be several minutes, so go grab a ☕!

_Once it has compiled, execute it_:

```bash
$ cd builddir  # builddir is default output directory
$ ./hello_world
Hello World!
```

### Packaging

The two relevant files in builddir are `hello_world` and a dynamically linked
library, `hello_world.$(dll)`. We can use a wildcard to tar both of them.

`tar -cvzf hello_world.tar.gz builddir/*`

To end result shoud be approximately 26MB, regardless of system.

## Conclusion

In this article, we constructed a Julia binary.

## See Also

### Further Reading

- [PackageCompiler](https://github.com/JuliaLang/PackageCompiler.jl)
- Other, similar articles
