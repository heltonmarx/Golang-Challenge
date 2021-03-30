# Transparent Cache Package

## Decisions taken and important notes

### 1. Check that the price was retrieved less than "maxAge" ago!

The solution adopted was the creating of an object `Price` with price float64 value
and also the timestamp of last use, updated at startup (`NewPrice`) and on each
valid retrieved.

A set of tests was also created to verify the methods related to the `Price` object.


###  2. Parallelize and optimized to not make the calls to the external service sequentially

The replacement of the `map[string]float64` by a [sync.Map](https://golang.org/pkg/sync/#Map) was adopted
to being able to access the products map concurrently.

The [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup) is used to wait for a collection of goroutines to finish,
and the [sync.Mutex](https://golang.org/pkg/sync/#Mutex) is used to mutual exclusion lock to append the results and errors
in the respective arrays.

In case of errors, empty array results was maintained, where [multierr](https://pkg.go.dev/go.uber.org/multierr) is used to
combine one or more errors together, returning the combination of all errors in only one.

A github workflow was included the execute the all tests with `-race` condition enabled and also examinate the source code with
suspicious constructs using [vet](https://golang.org/pkg/cmd/vet/).
