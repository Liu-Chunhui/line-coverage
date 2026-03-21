# line-coverage

A Go CLI tool that reads Go coverage profiles (`coverage.out`) and calculates line-by-line coverage metrics per file, with an overall percentage summary.

## Tech Stack
- Go 1.19
- CLI framework: `github.com/urfave/cli/v2`
- Logging: `github.com/sirupsen/logrus`
- Testing: `github.com/stretchr/testify`

## Project Structure
```
main.go                  # CLI entry point, flag definitions
cmd/report/report.go     # Report orchestration
pkg/coverage/            # Core coverage calculation logic
pkg/fileparser/          # File reading and regex patterns
pkg/percentage/          # Coverage percentage formatting
test/                    # Shared test utilities
```

## Common Commands
```bash
go build ./...           # Build
go test ./...            # Run all tests
go test ./... -v         # Verbose test output
```

## Language
All content written to this repository must be in English: code comments, commit messages, error messages, documentation, and any other text added to files.

## Coding Conventions
- All errors must be returned up the call stack — never use `panic`
- Use `logrus` for logging, never `fmt.Println` for debug output
- New CLI flags go in `main.go` under `app.Flags`; pass values into `report.Report()`
- Keep packages focused: coverage logic stays in `pkg/coverage/`, not in `cmd/`
- Test files live next to source files (`*_test.go` in the same package)

## Key Data Flow
```
coverage.out → profile.go (parse) → codebranch.go (line adjust) → targetresult.go (aggregate) → report.go (output)
```

## Module Name
`github.com/Liu-Chunhui/line-coverage`
