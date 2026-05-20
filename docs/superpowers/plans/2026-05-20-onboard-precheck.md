# Onboard Precheck — protoc-gen-checker

## Goal

Re-runnable diagnostic. Run before and after correction plans to verify project
state against the baseline captured during onboarding on 2026-05-20.

## Baseline (frozen, 2026-05-20)

- `go version` → `go1.26.2 linux/amd64`
- `go.mod` toolchain directive → `go 1.17`
- `golangci-lint --version` → `2.10.1`
- `golangci-lint run ./...` → exit 0, "No issues found" (default config, no `.golangci.yml`)
- `govulncheck -version` → `v1.2.0`
- `govulncheck ./...` → 0 reachable vulnerabilities; 3 in imported packages,
  8 in transitive modules (none reachable)
- `go test ./... -count=1 -short` → "no tests found", exit 0
- `go build ./...` → exit 0
- `.golangci.yml` → absent
- `.gitlab-ci.yml` / `.github/workflows/` → absent
- `vendor/` → absent
- `docs/DEVELOPMENT.md` → absent
- `*_test.go` → absent
- `AGENTS.md` → present (written by onboard, 2026-05-20)

## Tasks

### Task 1: Verify tooling installed
- [ ] `golangci-lint --version` exits 0
- [ ] `govulncheck -version` exits 0
- [ ] `go version` exits 0 and reports `go1.26.x` or newer

### Task 2: Run all mandatory checks
- [ ] `golangci-lint run ./...` — record exit code + issue count; compare to baseline
- [ ] `govulncheck ./...` — record reachable count + indirect counts; compare to baseline
- [ ] `go test ./... -count=1 -race` — record pass/fail; compare to baseline
- [ ] `go build ./...` — must exit 0
- [ ] `gofmt -l .` — must produce no output

### Task 3: Report deltas
- [ ] Append run result + delta vs baseline to this plan file under `## Run history`.
- [ ] If any check regressed, open a follow-up plan referencing the regression.

## Run history

### 2026-05-20 (baseline lock)
- golangci-lint: 0 issues (default config, Δ 0 vs baseline)
- govulncheck: 0 reachable, 3 imported / 8 module-level non-reachable (Δ 0)
- go test: no tests found (Δ 0)
- go build: ok
- gofmt -l .: `checker.go` was dirty — fixed in-flight (import grouping). Re-checked clean.
- Notes: ran during onboard execution pass on 2026-05-20. Locks baseline before correction plans.

```
### YYYY-MM-DD
- golangci-lint: <issues> (Δ <±N>)
- govulncheck: <reachable / indirect> (Δ <±N>)
- go test: <pass/fail/skipped> (Δ <±N>)
- go build: <ok/fail>
- gofmt -l .: <clean/dirty>
- Notes: ...
```
