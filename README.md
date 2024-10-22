# hotstash
A thread-safe key-value store for hot AWS lambda invocations that share execution environment.

Useful as a read cache or to deduplicate expensive db writes. Anything that fits into a key-value store really.

# Usage

- in-mem
- fs

# Test

````sh
$ go test
````

# TODOs
- [ ] comparable types for input
