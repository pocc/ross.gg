---
title: "Julia Versions"
pubDatetime: 2019-03-11T13:31:09Z
description: "Which version of Julia should you install? A guide to version differences."
tags:
  - julia
  - versions
---

_Which version of Julia should you install?_

## Table of contents

## >=1.0.0

Julia jumped from v0.6 to v0.7/v1.0.0 on 8 August 2018, which created breaking changes. On the upside, there are _significant_ [mathematical optimizations](https://discourse.julialang.org/t/fantastic-progress-in-master-branch/6868/2).

## <1.0.0

There are many libraries that depend on older Julia versions. Attempting to install one that hasn't been updated past v0.6 results in an "unsatisfiable requirements" error:

![Pkg.add failing due to Julia version incompatibility](@/assets/images/julia-versions/error_adding_playgrounds.webp)

If you are using an affected library:

1. Try to migrate to different >=v0.7 libraries with the same functionality
2. Create or add to existing issue on the repo regarding >=0.7 plans
3. Install a Julia Venv for v0.6 for this project

## Further Reading

- [Julia roadmap](https://github.com/JuliaLang/julia/milestones)
- [Project history](https://github.com/JuliaLang/julia/blob/master/HISTORY.md)
