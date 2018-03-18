setlocal enableDelayedExpansion
set PATH=%PRGS%\graphviz\latest\bin;%PATH%
set GOPATH=%~dp0..\..
go tool pprof -http=:8080 cpu.pprof
