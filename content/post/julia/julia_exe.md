---
title: "Making a Julia Binary"
date: 2019-03-04T16:20:52Z
author: Ross Jacobs
desc: "Make a Julia Binary using PackageCompiler"
keywords: binary
tags: 
  - draft1
image: "https://dl.dropboxusercontent.com/s/i7hlnqfd5lek700/julia_purple_exe.webp"

draft: true
---

# Making a Julia Binary
_Compile your Julia project with PackageCompiler for portability and speed._

Julia is a JIT language; however, sometimes it might be nice to have an
executable. There are two ways to interact with PackageCompiler: Using the CLI
using `juliac` and by adding PackageCompiler as part of your script. This
article will go over both.

## Initial Setup

1. [Install Julia 1.1.0](https://julialang.org/downloads/) if not installed
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

_Use your shell of choice below_

#### bash
   
```bash
$ juliac_path='julia -e println(normpath(Base.find_package(\
	"PackageCompiler"),"..", "..")'
$ echo "julia ${juliac_path}juliac.jl" >> ~/.bashrc
$ source ~/.bashrc
```

#### powershell
	
```powershell
PS> $juliac_dir = julia -e 'println(normpath(Base.find_package(`
		\"PackageCompiler\"),\"..\",\"..\"))'
PS> Set-Alias julia "${juliac_dir}juliac.jl" 
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
executable. I will cover options here briefly for sake of completeness
(use `juliac -h` for full details).

- __-t__ Remove temp build files 
- __-a__ Automatically build required dependencies
- __-v__ Verbose
- __-e__ Make executable
- __-r__ Implies -O3 -g0 [Max (-O)ptimize and debu(-g) at the lowest level]
- __-n__ Output file name

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

In this article, we constructed a Julia binary In my [second article](/tbd) on the
topic, I explore how to encapsulate Julia's Calculus library into a CLI utility.

## See Also

### Further Reading

- [PackageCompiler](https://github.com/JuliaLang/PackageCompiler.jl)
- Other, similar articles

### 


<img
  src="http://www.quickmeme.com/img/92/927d52fd29f08027c5356e5f8bfd78021dcd2351d18d717eb86d393132f7322a.jpg"
  alt="FIXME"
  style="width:100%;height=100%;text-align:left"/>

## Further Reading

**Questions and Exercises**

- Question/Exercise
- Question/Exercise

**Relevant Articles**

- [Article 1]()
- [Article 2]()

**Sources** [0]() [1]()

## Drafting

### Prewriting

**Audience**

Who is your audience?

**Deliverable**

What is the ONE thing your audince gain from reading this?

**Niche**

What makes this unique compared to existing articles? 

### Checklist

** Basic**
* [ ] Intro: How WILL they get the deliverable?
* [ ] 300-600 words
* [ ] Images: Cover image, Reengage image/table
* [ ] Conclusion: How DID they get the deliverable?
* [ ] Questions/Exercises/Call To Action

**Extended**
* [ ] Keywords: Front Matter, Title, Desc, Post: (top, end), Images: (alt, title)
* [ ] 3-4 external links
* [ ] 1-2 sources
* [ ] 2-4 internal links
* [ ] Lint!

### Prepublish

- Engaging: Why will the reader read until the end?
- Organized: Identify specific things that the reader might be looking
  for in subsections. How easy are they to find?
- Optimized: Can the Deliverable be provided to the reader in fewer words?

