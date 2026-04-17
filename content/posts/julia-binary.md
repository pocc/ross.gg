---
title: "Making a Julia Binary"
date: 2019-03-15T16:20:52Z
description: "Compile your Julia project into a portable executable using PackageCompiler."
tags:
  - julia
  - compilation
stage: "evergreen"
related:
  - "/notes/installing-julia"
  - "/notes/julia-versions"
draft: false
---

![](/img/julia/julia_purple_exe.webp)

*Compile your Julia project with PackageCompiler for portability*

Julia is a JIT language; however, sometimes it might be nice to have an executable. There are two ways to interact with PackageCompiler: Using the CLI using `juliac` and by adding PackageCompiler as part of your script. This article will go over both.

## Initial Setup

1. [Install Julia 1.1.0](https://julialang.org/downloads/) if absent
2. Install PackageCompiler and Deps

```julia
julia -e 'using Pkg;
      Pkg.add("ArgParse");
      Pkg.add("PackageCompiler")'
```

### Alias PackageCompiler to `juliac`

The full command to compile is `julia /variable/path/to/juliac.jl`. The path to this file is dependent on where PackageCompiler is stored and different on every system. We are going to save time by using an alias.

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

```julia
# hello_world.jl
function AwesomeFunction()
  println("Hello World!")
end

Base.@ccallable function julia_main(ARGS::Vector{String})::Cint
    AwesomeFunction()
    return 0
end
```

### Compiling

Use `juliac hello_world.jl -tavern hello_world` to generate the executable:

- **-t** Remove temp build files
- **-a** Automatically build required dependencies
- **-v** Verbose
- **-e** Make executable
- **-r** Implies -O3 -g0
- **-n** Output file name

```bash
$ cd builddir
$ ./hello_world
Hello World!
```

### Packaging

```bash
tar -cvzf hello_world.tar.gz builddir/*
```

The result should be approximately 26MB, regardless of system.

## See Also

- [PackageCompiler](https://github.com/JuliaLang/PackageCompiler.jl)
