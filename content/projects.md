---
title: "Projects"
description: "Things completed"
date: 2019-03-01T17:17:03Z
author: Ross Jacobs
draft: false
---

## [tshark.dev](https://github.com/pocc/tshark.dev)

![](/img/projects/tshark.dev.png)

I didn't like the documentation for Wireshark's CLI, so I completely documented all Wireshark command line utilities. 



## [Pre-commit-hooks](https://github.com/pocc/pre-commit-hooks)

![](/img/projects/pre-commit-hooks.png)

[pre-commit](https://pre-commit.com) hooks repo that
integrates two C/C++ code formatters:
> [clang-format](https://clang.llvm.org/docs/ClangFormatStyleOptions.html),
[uncrustify](http://uncrustify.sourceforge.net/),

and five C/C++ static code analyzers:
> [clang-tidy](https://clang.llvm.org/extra/clang-tidy/),
[oclint](http://oclint.org/),
[cppcheck](http://cppcheck.sourceforge.net/),
[cpplint](https://github.com/cpplint/cpplint),
[include-what-you-use](https://github.com/include-what-you-use/include-what-you-use)

### Additional features

* Relay correct pass/fail to pre-commit, even when some commands exit 0 when they should not. Some versions of oclint, clang-tidy, and cppcheck have this behavior.
* Honor `--` arguments, which pre-commit [has problems with](https://github.com/pre-commit/pre-commit/issues/1000)
* Optionally [enforce a command version](https://github.com/pocc/pre-commit-hooks#special-flags-in-this-repo) so your team gets code formatted/analyzed the same way
* Formatters clang-format and uncrustify will error with diffs of what has changed