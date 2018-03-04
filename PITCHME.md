# Go Profilling

Performance profiling

Note:

- [Go Profilling](#go-profilling)
  - [Plan](#plan)
    - [About this presentation](#about-this-presentation)
    - [About me (Daniel)](#about-me-daniel)
    - [About me (VonC)](#about-me-vonc)
  - [Why Profiling](#why-profiling)
    - [For reporting](#for-reporting)
    - [For testing](#for-testing)
    - [For Measuring](#for-measuring)
  - [Performance profiling](#performance-profiling)
    - [Example: Julia Set](#example-julia-set)
    - [Event-based](#event-based)
    - [setup](#setup)
    - [CPU](#cpu)
    - [Memory](#memory)
    - [Problem](#problem)
  - [Event-based profiling](#event-based-profiling)
    - [Tracer](#tracer)
    - [Goroutine vs. GC](#goroutine-vs-gc)
    - [Trade-off](#trade-off)
  - [Event-based Profiling](#event-based-profiling)
  - [Steps](#steps)
    - [Add Profiling](#add-profiling)
    - [What do we see](#what-do-we-see)
  - [Benchmark](#benchmark)

---
<!-- .slide: data-background="#030202" -->

## Plan

- Why Profiling
- What: Performance Profiling
- How: Event-based Profiling

+++

### About this presentation

- Available at [github.com](https://github.com/VonC/talks/blob/2018_goprofiling/PITCHME.md)</a>
- Available at [gitpitch.com](https://gitpitch.com/VonC/talks/2018_goprofiling?grs=github)</a>
- Available at intranet.softeam.com:  
  "[Go Fractal, no bugs!!](https://intranet.softeam.fr/node/2904)" (Oct. 2017)
- Fully annotated

Note:

You will find in the shownotes, or directly in the markdown article on GitHub
additional information with each slides.

Palette: <http://paletton.com/#uid=1000u0k004h0jin01bM5n02dm0p>

+++

### About me (Daniel)

![Chaffiol](assets/img/chaffiol.jpg)

**Daniel CHAFFIOL** (Softeam)

- Since 1999
- Development architect
- BNP, SGCIB, HSBC, Amundi

Note:

- Full CV: <https://stackoverflow.com/cv/vonc>

+++

### About me (VonC)

![VonC](assets/img/vonc.jpg)

**VonC** (Stack Overflow)

- Since 2008
- 4th all-time user
- Topics: Version Control (Git), Go, Docker

Note:

- Stack Overflow profile: <https://stackoverflow.com/users/6309/vonc>

---
<!-- .slide: data-background="#030202" -->

## Why Profiling

- For Reporting
- For Testing
- For Measuring

Note:

What is profiling?

3 kinds of profiling ("Performance Profiling" <http://thomas-solignac.com/blog/slides-talk-05-11-2017-performance-profiling/>
from Thomas Solignac <https://twitter.com/thomassolignac?lang=en>)

+++

### For reporting

- Services continuous monitoring
  - Availability
  - Latency
  - EventLog

+++

### For testing

- Code Profiling
  - Dependencies
  - Code qualities (linters)
  - Call Graph

Note:

Testing techniques are numerous with Go: <https://speakerdeck.com/mitchellh/advanced-testing-with-go>
(Mitchell Hashimoto: <https://twitter.com/mitchellh>)

+++

### For Measuring

- Perfomance profiling
  - CPU
  - Memory

---

## Performance profiling

APM:

- Statistical
- Event-based

Note:

2 big categories within APM Application Performance Management (<https://en.wikipedia.org/wiki/Application_performance_management>)

<https://en.wikipedia.org/wiki/Profiling_(computer_programming)>

<https://www.raymond.cc/blog/measure-time-taken-to-complete-a-batch-file-or-command-line-execution/>
<https://code.google.com/archive/p/time-windows/source/default/source>
<https://github.com/golang/benchmarks/blob/master/driver/driver_windows.go>

+++

### Example: Julia Set

+++

### Event-based

### setup

- go dep
- graphviz

Note:

#### Graph

For the graphic GUI version of profiling, You will need:

- "**Graphviz - Graph Visualization Software**" (<https://graphviz.gitlab.io>)  
  Windows Packages: <https://graphviz.gitlab.io/_pages/Download/Download_windows.html>

#### Dependencies

Project **golang/dep** (<https://github.com/golang/dep>)
from **Sam Boyer** (<https://twitter.com/sdboyer>).  
See "**So you want to write a package manager**" (<https://medium.com/@sdboyer/so-you-want-to-write-a-package-manager-4ae9c17d9527#.740o43vxi>)

```bash
dep init
dep ensure
dep status -v
```

+++

### CPU

+++

### Memory

+++

### Problem

---

## Event-based profiling

+++

### Tracer

+++

### Goroutine vs. GC

+++

### Trade-off

Note:

But it does not stop here:

- Code coverage: package by package, or (with the Go 1.10, for the all project: <https://github.com/golang/go/issues/16768>)
- See "Building and using coverage-instrumented programs with Go (<http://damien.lespiau.name/2017/05/building-and-using-coverage.html>)
  from DAMIEN LESPIAU (<https://twitter.com/__damien__>)
- Profiling: it is even better with the Go 1.10: "The new pprof user interface"
  (<https://rakyll.org/pprof-ui/>) from Jaana Burcu Dogan (JBD), aka @rakyll (<https://twitter.com/rakyll>)

Profiler pprof "**Go Profiler Internals**" (<https://stackimpact.com/blog/go-profiler-internals/>)
from **Dmitri Melikyan** (<https://github.com/dmelikyan>)
founder of **StackImpact** (<https://twitter.com/stackimpact>).

Main example: "**Using the Go execution tracer to speed up fractal rendering**" (<https://medium.com/@francesc/using-the-go-execution-tracer-to-speed-up-fractal-rendering-c06bb3760507>,
<https://campoy.cat/blog/using-the-go-tracer-to-speed-up-fractal-making/>)
from **Francesc Campoy** (<https://twitter.com/francesc>):

- His code is at <https://github.com/campoy/justforfunc/tree/master/22-perf>
- And a video demonstration is available at <https://www.youtube.com/watch?v=ySy3sR1LFCQ&feature=youtu.be&list=PL6>

Based on code from <https://github.com/sfluor/fractcli>, authored
by **Salph Tabet** (<https://github.com/sfluor>)  
It uses also <https://github.com/lucasb-eyer/go-colorful>, a library for playing
with colors in go (golang), by **Lucas Beyer** (<http://lucasb.eyer.be/>).

Cf. "**Understanding Julia and Mandelbrot Sets**" (<http://www.karlsims.com/julia.html>)
by **Karl Sims** (<http://www.karlsims.com/>)

See also "**Profiling Go**" <http://www.integralist.co.uk/posts/profiling-go/>
from **Mark McDonnell** (<https://twitter.com/integralist>)

## Event-based Profiling

New pprof UI with Go 1.10
See "**The new pprof user interface**" <https://rakyll.org/pprof-ui/> from rakyll

See also <https://medium.com/@cep21/using-go-1-10-new-trace-features-to-debug-an-integration-test-1dc39e4e812d>

    go tool pprof -http=:8080 cpu.pprof

It is a bit of an hassle to trigger the profiling, redirecting its output to a file
(see <https://groups.google.com/forum/#!topic/golang-nuts/YhnyJDI3IG0>).  
But you have "**pkg/profile**" (<https://github.com/pkg/profile>) from **Dave Cheney**
(<https://github.com/davecheney>, <https://dave.cheney.net/>, <https://twitter.com/davecheney>)

See "**PROFILING GO APPLICATIONS WITH FLAMEGRAPHS**" (<http://brendanjryan.com/golang/profiling/2018/02/28/profiling-go-applications.html>)"
from **Brendan Ryan** (<https://twitter.com/brendan_j_ryan>) for the Uber approach
to flamegraph.

But this approach has now been superseded with the alternative `pprof` tool,
with flamegraph support:

<https://github.com/google/pprof>

    pprof.exe -http=:8080 cpu.pprof

Tracer (hooks)

## Steps

### Add Profiling

### What do we see

With pprf, only what is executed, each time we are asking the program.

Abs.Cplx: <http://agniva.me/go/2017/08/27/fun-with-go-assembly.html>

## Benchmark

<https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go>

    go get -u golang.org/x/tools/cmd/benchcmp
    go test -bench=Simple
    go test -bench=PerPixel
    go test -bench=PerCol

Note: time is monotonic since Go 1.9: <https://github.com/golang/proposal/blob/master/design/12914-monotonic.md>

    100          31303975 ns/op

- <https://github.com/golang/proposal/blob/master/design/14313-benchmark-format.md>
- <https://github.com/cespare/prettybench>

````(bash)
> go test -bench=.
goos: windows
goarch: amd64
pkg: julia_raw
Benchmark_createImageSimple-4                 20          96612270 ns/op
Benchmark_createImageGoPerPixel-4             20          99012575 ns/op
Benchmark_createImageGoPerCol-4              100          13116665 ns/op
PASS
ok      julia_raw       7.462s
````