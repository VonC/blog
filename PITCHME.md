# Go Profilling

Note:

Testing techniques are numerous with Go: <https://speakerdeck.com/mitchellh/advanced-testing-with-go>
(Mitchell Hashimoto: <https://twitter.com/mitchellh>)  
But it does not stop here:

- Code coverage: package by package, or (with the Go 1.10, for the all project: <https://github.com/golang/go/issues/16768>)
- See "Building and using coverage-instrumented programs with Go (<http://damien.lespiau.name/2017/05/building-and-using-coverage.html>)
  from DAMIEN LESPIAU (<https://twitter.com/__damien__>)
- Profiling: it is even better with the Go 1.10: "The new pprof user interface"
  (<https://rakyll.org/pprof-ui/>) from Jaana Burcu Dogan (JBD), aka @rakyll (<https://twitter.com/rakyll>)

But what is profiling?

3 kinds of profiling ("Performance Profiling" <http://thomas-solignac.com/blog/slides-talk-05-11-2017-performance-profiling/>
from Thomas Solignac <https://twitter.com/thomassolignac?lang=en>)

- Code Profiling
  - Dependencies
  - Code qualities (linters)
  - Call Graph
- Services continuous monitoring
  - Availability
  - Latency
  - EventLog
- Perfomance profiling
  - CPU
  - Memory

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

## Performance Profiling

2 big categories within APM Application Performance Management (<https://en.wikipedia.org/wiki/Application_performance_management>)

- Statistical
- Event-based

<https://en.wikipedia.org/wiki/Profiling_(computer_programming)>

<https://www.raymond.cc/blog/measure-time-taken-to-complete-a-batch-file-or-command-line-execution/>
<https://code.google.com/archive/p/time-windows/source/default/source>
<https://github.com/golang/benchmarks/blob/master/driver/driver_windows.go>

### Statistical

pprof

### Event-based

New pprof UI with Go 1.10
See "**The new pprof user interface**" <https://rakyll.org/pprof-ui/> from rakyll

It is a bit of an hassle to trigger the profiling, redirecting its output to a file
(see <https://groups.google.com/forum/#!topic/golang-nuts/YhnyJDI3IG0>).  
But you have "**pkg/profile**" (<https://github.com/pkg/profile>) from **Dave Cheney**
(<https://github.com/davecheney>, <https://dave.cheney.net/>, <https://twitter.com/davecheney>)

Tracer (hooks)

## Steps

### Dependencies

Project **golang/dep** (<https://github.com/golang/dep>)
from **Sam Boyer** (<https://twitter.com/sdboyer>).  
See "**So you want to write a package manager**" (<https://medium.com/@sdboyer/so-you-want-to-write-a-package-manager-4ae9c17d9527#.740o43vxi>)

```bash
dep init
dep ensure
dep status -v
```

### Add Profiling

For the graphic GUI version of profiling, You will need:

- "**Graphviz - Graph Visualization Software**" (<https://graphviz.gitlab.io>)  
  Windows Packages: <https://graphviz.gitlab.io/_pages/Download/Download_windows.html>

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