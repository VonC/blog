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

Profiler pprof "Go Profiler Internals" (<https://stackimpact.com/blog/go-profiler-internals/>)
from Dmitri Melikyan (<https://github.com/dmelikyan>) founder of StackImpact (<https://twitter.com/stackimpact>).

Main example: "Using the Go execution tracer to speed up fractal rendering" (<https://medium.com/@francesc/using-the-go-execution-tracer-to-speed-up-fractal-rendering-c06bb3760507>,
<https://campoy.cat/blog/using-the-go-tracer-to-speed-up-fractal-making/>)
from Francesc Campoy (<https://twitter.com/francesc>):

- His code is at <https://github.com/campoy/justforfunc/tree/master/22-perf>
- And a video demonstration is available at <https://www.youtube.com/watch?v=ySy3sR1LFCQ&feature=youtu.be&list=PL6>

Based on code from <https://github.com/sfluor/fractcli>, authored
by Salph Tabet (<https://github.com/sfluor>)  
It uses also <https://github.com/lucasb-eyer/go-colorful>, a library for playing
with colors in go (golang), by Lucas Beyer (<http://lucasb.eyer.be/>).