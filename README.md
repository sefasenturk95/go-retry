# go-retry

A simple Go package for retrying operations with optional exponential backoff.

## Installation

```bash
go get github.com/sefasenturk95/go-retry
```

## Usage

### Basic Retry

```go
import "github.com/sefasenturk95/go-retry"

err := retry.Do(3, func() error {
    return someOperation()
})
```

### Retry with Exponential Backoff

```go
// Retries up to 3 times with delays: 1s, 2s, 4s
err := retry.WithBackOff(3, 1000, func() error {
    return someOperation()
})
```

The backoff delay doubles each attempt:
- Attempt 1: no delay
- Attempt 2: 1s delay
- Attempt 3: 2s delay
- Attempt 4: 4s delay
- etc.

## Error Handling

Errors are wrapped with `%w`, so you can use `errors.Is()` and `errors.As()`:

```go
err := retry.Do(3, func() error {
    return ErrNotFound
})

if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

## License

MIT
