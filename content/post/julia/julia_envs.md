---
title: "Julia Envs"
date: 2019-03-11T13:31:09Z
author: Ross Jacobs
desc: "Using a specific Julia environment for your project"
tags:
  - julia
  - environment
image: https://www.residentadvisor.net/images/clubs/versions_logo_ra_312x210px.png

draft: true
---

_Using a specific Julia environment for your project_

Note: You should check out [Julia Environment
Variables](https://docs.julialang.org/en/v1/manual/environment-variables/index.html)
if you haven't already.

## Julia ENVs

There are a couple ways to structure your Julia project. In this example, we
want to install `Calculus` and `Polynomials` on `Julia==0.6`.

### Set up julia and packages for this project

1. [Set up a project]()
2. In project dir, start `julia`
3. You can enter pkg mode with `]` in the Julia interpreter. 
   enter in pkg mode, you can see all of the available commands. If any of the
   following commands do not make sense, I encourage you to return to pkg help. 
4. Activate the directory (Similar to python venv's `source venv/bin/activate`)
    
	{{< highlight julia >}}
    julia> using Pkg
	julia> Pkg.activate("."){{< /highlight >}}

5. Change the delevop dir for this Julia shell. Per the
   [documentation](https://docs.julialang.org/en/v1/stdlib/Pkg/index.html),  The
   develop dir is where libraries are downloaded so that you can look at the
   code locally.  
    
	{{< highlight julia >}}julia> ENV["JULIA_PKG_DIR"] ="$(pwd())/venv"{{< /highlight >}}

6. Install your libraries locally

	{{< highlight julia >}}
	deps = ["Calculus", "Polynomials"]
	julia> Pkg.develop(deps){{< /highlight >}}
    
### Setting up a direnv

* [ ] Talk about https://direnv.net/
* [ ] Setup ENV["JULIA_BINDIR"] (see
	  https://docs.julialang.org/en/v1/manual/environment-variables/index.html)

## See Also
* [Playground.jl](https://github.com/rofinn/Playground.jl): A way to create
  Julia venvs, but currently held back at Julia v0.6.
* [julia-venv](https://github.com/tkf/julia-venv): Julia venvs _inside of Python_.
* [venv.jl](https://github.com/kdmurray91/venv.jl), julia venv copied from pyenv. Last
  commit was in 2016.

