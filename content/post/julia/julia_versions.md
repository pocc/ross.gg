---
title: "Julia Versions"
date: 2019-03-11T13:31:09Z
author: Ross Jacobs
desc: "Different Julia versions"
tags:
  - julia
  - versions
image: https://www.residentadvisor.net/images/clubs/versions_logo_ra_312x210px.png

draft: false
---

_Which version of Julia should you install?_

## Minor Differences

If at all possible, you should [install Julia](/post/installing-julia) for latest
versions of Julia.

### >=1.0.0

Julia jumped from v0.6 to v0.7/v1.0.0 on 8 August 2018, which created breaking
changes. On the upside, there are _significant_ [mathematical
optimizations](https://discourse.julialang.org/t/fantastic-progress-in-master-branch/6868/2).

### <1.0.0

There are many libraries that depend on older Julia versions. For example,
[Mocha](https://github.com/pluskid/Mocha.jl) (ML Framework),
[FemtoCleaner](https://github.com/JuliaComputing/FemtoCleaner.jl) (Upgrades
deprecated syntax), [ACME](https://github.com/HSU-ANT/ACME.jl) (Circuit
Modeling), and [Playground](https://github.com/rofinn/Playground.jl) (Julia
VENVs) are all tied to v0.6. If you are using an affected libraries, you should
do the following:

1. Try to migrate to different >=v0.7 libraries with the same functionality
2. Create or add to existing issue on the repo regarding >=0.7 plans
3. Install a Julia Venv <!-- Link Me when done /post/julia-envs --> for v0.6 for this project

Keep in mind that projects get upgraded, so it's worthwhile to occasionally
check back on repos to see if they've fixed it.  

## Further Reading

* Check out the Julia [roadmap](https://github.com/JuliaLang/julia/milestones)
  to see what features are planned! 
* The [project
  history](https://github.com/JuliaLang/julia/blob/master/HISTORY.md) is good if
  you want to know more about how releases differ.

