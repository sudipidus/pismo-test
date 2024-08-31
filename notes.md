1. Adding pprof:
importing net/http/pprof works with default serve mux so we have to declare handler
apart from import

import _ "net/http/pprof"
r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)

(notes: should this be exposed on prod publicly? no)

go tool pprof http://localhost:8080/debug/pprof/heap
(use commands like top)

creating png (visualization)

go tool pprof -png http://localhost:8080/debug/pprof/heap > out.png




2. zero allocation loggers
what is zero allocation
running go benchmark tests



STRACE:
strace is a utility to see system calls made by a running process
helpful for debugging and understanding what's happening internally

eg: strace ls
```sh
openat(AT_FDCWD, "/lib/x86_64-linux-gnu/libselinux.so.1", O_RDONLY|O_CLOEXEC) = 3
read(3, "\177ELF\2\1\1\0\0\0\0\0\0\0\0\0\3\0>\0\1\0\0\0\0\0\0\0\0\0\0\0"..., 832) = 832
newfstatat(3, "", {st_mode=S_IFREG|0644, st_size=166280, ...}, AT_EMPTY_PATH) = 0
mmap(NULL, 177672, PROT_READ, MAP_PRIVATE|MAP_DENYWRITE, 3, 0) = 0x7f353e2e5000
mprotect(0x7f353e2eb000, 139264, PROT_NONE) = 0
mmap(0x7f353e2eb000, 106496, PROT_READ|PROT_EXEC, MAP_PRIVATE|MAP_FIXED|MAP_DENYWRITE, 3, 0x6000) = 0x7f353e2eb000

```



### Escape analysis in go:
compile time optimization

stack is faster/expensive than heap
but if a function's variable has to outlive the scope (should be retained
even after function is execute, eg: say it's returning a pointer), then such variables
escape to heap

`
go build -gcflags '-m' .
`
(to see what escapes to heap)



### golangci-lint

```azure
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3

golangci-lint --version

golangcli-lint run ./...
```


#### creating new migration
migrate create -ext sql -dir db/migrations -seq change_id_to_int_in_operation_types  