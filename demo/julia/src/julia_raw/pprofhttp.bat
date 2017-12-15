setlocal enableDelayedExpansion
set PATH=%PRGS%\graphviz\latest\release\bin;%PATH%
set GOPATH=%~dp0..\..
go tool pprof -http=:8080 %~dp0..\..\bin\cpu.pprof
