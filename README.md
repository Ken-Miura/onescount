# onescount
## go test -bench .
goos: linux  
goarch: amd64  
pkg: github.com/Ken-Miura/onescount  
BenchmarkOnesCount-8         	20000000	        73.7 ns/op  
BenchmarkStdLibOnesCount-8   	2000000000	         0.38 ns/op  
PASS  
ok  	github.com/Ken-Miura/onescount	2.354s  

## cat /proc/cpuinfo | grep "model name"
model name	: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz

## cat /proc/meminfo | grep "MemTotal"
MemTotal:       32888328 kB
