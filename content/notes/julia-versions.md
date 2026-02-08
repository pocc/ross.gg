---
title: "Julia Versions"
date: 2019-03-11T13:31:09Z
description: "Which version of Julia should you install? A guide to version differences."
tags:
  - julia
  - versions
stage: "evergreen"
related:
  - "/notes/installing-julia"
  - "/notes/julia-binary"
draft: false
---

*Which version of Julia should you install?*

## >=1.0.0

Julia jumped from v0.6 to v0.7/v1.0.0 on 8 August 2018, which created breaking changes. On the upside, there are *significant* [mathematical optimizations](https://discourse.julialang.org/t/fantastic-progress-in-master-branch/6868/2).

## <1.0.0

There are many libraries that depend on older Julia versions. If you are using an affected library:

1. Try to migrate to different >=v0.7 libraries with the same functionality
2. Create or add to existing issue on the repo regarding >=0.7 plans
3. Install a Julia Venv for v0.6 for this project

## Further Reading

* [Julia roadmap](https://github.com/JuliaLang/julia/milestones)
* [Project history](https://github.com/JuliaLang/julia/blob/master/HISTORY.md)
