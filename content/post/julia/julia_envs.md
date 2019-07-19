---
title: "Julia Envs"
date: 2019-07-19T13:31:09Z
author: Ross Jacobs
description: "Julia Environments for v0.6+"
tags:
  - julia
  - virtualenv
image: https://www.residentadvisor.net/images/clubs/versions_logo_ra_312x210px.png

draft: true
---

Create virtualenvs for Julia v0.6+

[Playgrounds](https://github.com/rofinn/Playground.jl) is a good way to create Julia Envs. However, it is limited to Julia v0.5, which is a deprecated version.
Once Playgrounds supports Julia for latest Julia, this article will no longer be relevant.

[Relevant Discussion](https://discourse.julialang.org/t/handling-multiple-versions-of-julia/14035), with people trying to solve
the problem of multiple Julias.

![](https://dl.dropboxusercontent.com/s/4dxhudowee98n4e/error_installing_playgrounds.png)

The current alternative is to use [direnv](https://direnv.net/) and shell variables,
which is what this article demonstrates.
If you have not checked out [Julia Environment
Variables](https://docs.julialang.org/en/v1/manual/environment-variables/index.html) yet,
you probably should.

**NOTE**: Direnv is only available for Unix systems, so if you are on Windows, use WSL.

## Julia ENVs

There are a couple ways to structure your Julia project. In this example, we
arbitrarily want to install `Calculus` and `Polynomials` on `Julia==0.6`.

### Set up julia and packages for this project

1. Create the projcet

  ```bash
  mkdir ExampleProject
  cd ExampleProject
  # Create the virtualenv folder where libraries are saved. 'venv' is arbitrary.
  mkdir venv
  # Create
  cat << EOF > project.toml
  name = "ExampleProject"
  uuid = "$(uuid)"
  version = "0.0.0"

  [compat]
  julia = "1.1"
  EOF
  ```

1. In project dir, start `julia`
1. You can enter pkg mode with `]` in the Julia interpreter.
   enter in pkg mode, you can see all of the available commands. If any of the
   following commands do not make sense, I encourage you to return to pkg help.

1. Activate the directory (Similar to python venv's `source venv/bin/activate`)
  ```julia
  julia> using Pkg
  julia> Pkg.activate(".")
  ```

1. Change the delevop dir for this Julia shell. Per the
   [documentation](https://docs.julialang.org/en/v1/stdlib/Pkg/index.html),  The
   develop dir is where libraries are downloaded so that you can look at the
   code locally.

   ```julia
   ENV["JULIA_PKG_DEVDIR"] = "$(pwd())/venv"
   ```

   More information on this variable can be found in [this discussion](https://discourse.julialang.org/t/julia-pkgdir-and-julia-0-7/11672/6).

1. Install your libraries locally
  ```julia
  deps = ["Calculus", "Polynomials"]
  julia> Pkg.develop(deps)
  ```

1. (Optional) Updating your local env in a later Julia shell. You can use
  [Package Levels](https://julialang.github.io/Pkg.jl/v1/api/#Pkg.UpgradeLevel)
  if you have a specific update in mind.
  ```julia
  julia> ENV["JULIA_PKG_DEVDIR"] = "$(pwd())/venv"
  deps = ["List", "Of", "All", "My", "Libraries"]
  Pkg.update(deps)
  ```

### Setting up a direnv

* [ ] Need to figure out what default 
* [ ] Talk about https://direnv.net/
* [ ] Setup ENV["JULIA_BINDIR"] (see
	  https://docs.julialang.org/en/v1/manual/environment-variables/index.html)

## See Also

* [Pkg](https://docs.julialang.org/en/v1.0/stdlib/Pkg/): Official docs for using the Pkg utility.
* [Playground.jl](https://github.com/rofinn/Playground.jl): A way to create
  Julia venvs, but currently held back at Julia v0.6.
* [julia-venv](https://github.com/tkf/julia-venv): Julia venvs _inside of Python_.
* [venv.jl](https://github.com/kdmurray91/venv.jl): Julia venv proof of concept.
