---
title: "Using Urllib"
date: 2019-03-14T11:35:30Z
author: Ross Jacobs
desc: "Using Urllib"
tags: 
  - python
  - urllib
image: https://i.stack.imgur.com/wMYZG.png

draft: true
---

_Urllib is a part of the Python Standard Library_

[Urllib Docs](https://docs.python.org/3/library/urllib.html)

[Requests](http://docs.python-requests.org/en/master/) makes Python web scraping
easy. It also has an inspiring API, which is recommended reading for Python
neophytes. However, requests is not in the Python
standard library, and this will also not change any time soon [0] [1] [2]. I do not
think this is a bad thing though: Requests is an API around existing Python
libraries. If you want to use this easy to use API, that is your choice; however
you will be adding a dependency to your project.

In my opinion reducing external dependencies makes for cleaner code.
This article is aimed at people who are willing to use a clunkier http interface
(urllib) to reduce dependencies. This article provides examples of urllib and how to
migrate from requests to urllib.

[0]: [Asking the community for input](https://github.com/kennethreitz/requests/issues/2424)
<br>[1]: [Kenneth Reitz on stdlib inclusion](http://docs.python-requests.org/en/master/dev/philosophy/#standard-library)
<br>[2]: [LWN discussion based on [1]](https://lwn.net/Articles/640838/)

## Asserts
- You are a python programmer who interacts with websites 

## What reasons are there for sticking with Requsets? 
(urllib might be behind the times when it comes to ssl)

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
