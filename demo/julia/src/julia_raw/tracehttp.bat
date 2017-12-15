setlocal enableDelayedExpansion
set PATH=%PRGS%\graphviz\latest\release\bin;%PATH%
pprof -http=:8080 trace.out
