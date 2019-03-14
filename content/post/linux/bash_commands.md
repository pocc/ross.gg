---
title: "Bash "
date: 2019-03-11T19:35:02Z
author: Ross Jacobs
desc: "Popularity of various coreutils"
keywords: draft
tags:
  - bash
image: http://www.quickmeme.com/img/92/927d52fd29f08027c5356e5f8bfd78021dcd2351d18d717eb86d393132f7322a.jpg

draft: true
---

_Where do bash utilites come from?_

Typing in `help` or `bash -c help` will give you a list of all commands.
We can extract the commands themselves like so:

```bash
bash -c help \
| perl -ne 'print "$1\n" while /(?:^| ) ([-_\(\[\.a-{0-:]+(?: \[&\])?)/g' \
| sort
```

| Command      | What does it do? | Source |
|--------------|------------------|--------|
| ((           |                  |        |
| .            |                  |        |
| :            |                  |        |
| [            |                  |        |
| [[           |                  |        |
| alias        |                  |        |
| bg           |                  |        |
| bind         |                  |        |
| break        |                  |        |
| builtin      |                  |        |
| caller       |                  |        |
| case         |                  |        |
| cd           |                  |        |
| command      |                  |        |
| compgen      |                  |        |
| complete     |                  |        |
| compopt      |                  |        |
| continue     |                  |        |
| coproc       |                  |        |
| declare      |                  |        |
| dirs         |                  |        |
| disown       |                  |        |
| echo         |                  |        |
| enable       |                  |        |
| eval         |                  |        |
| exec         |                  |        |
| exit         |                  |        |
| export       |                  |        |
| false        |                  |        |
| fc           |                  |        |
| fg           |                  |        |
| for          |                  |        |
| for          |                  |        |
| function     |                  |        |
| getopts      |                  |        |
| hash         |                  |        |
| help         |                  |        |
| history      |                  |        |
| if           |                  |        |
| job_spec [&] |                  |        |
| jobs         |                  |        |
| kill         |                  |        |
| let          |                  |        |
| local        |                  |        |
| logout       |                  |        |
| mapfile      |                  |        |
| popd         |                  |        |
| printf       |                  |        |
| pushd        |                  |        |
| pwd          |                  |        |
| read         |                  |        |
| readarray    |                  |        |
| readonly     |                  |        |
| return       |                  |        |
| select       |                  |        |
| set          |                  |        |
| shift        |                  |        |
| shopt        |                  |        |
| source       |                  |        |
| suspend      |                  |        |
| test         |                  |        |
| time         |                  |        |
| times        |                  |        |
| trap         |                  |        |
| true         |                  |        |
| type         |                  |        |
| typeset      |                  |        |
| ulimit       |                  |        |
| umask        |                  |        |
| unalias      |                  |        |
| unset        |                  |        |
| until        |                  |        |
| variables    |                  |        |
| wait         |                  |        |
| while        |                  |        |
| {            |                  |        |


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

- [ ] Intro: How WILL they get the deliverable?
- [ ] 300-600 words
- [ ] Images: Cover image, Reengage image/table
- [ ] Conclusion: How DID they get the deliverable?
- [ ] Questions/Exercises/Call To Action

**Extended**

- [ ] Keywords: Front Matter, Title, Desc, Post: (top, end), Images: (alt,
      title)
- [ ] 3-4 external links
- [ ] 1-2 sources
- [ ] 2-4 internal links
- [ ] Lint!

### Prepublish

- Engaging: Why will the reader read until the end?
- Organized: Identify specific things that the reader might be looking for in
  subsections. How easy are they to find?
- Optimized: Can the Deliverable be provided to the reader in fewer words?
