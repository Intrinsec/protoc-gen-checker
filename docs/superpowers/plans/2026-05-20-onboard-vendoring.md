# Vendoring

## Goal

Vendor Go dependencies so builds are reproducible without network access and so
CI can run `-mod=vendor`. Mandatory for tier B Go projects.

## Context

Repo has no `vendor/` directory at onboard time. `go.mod` + `go.sum` are
authoritative; `go mod vendor` adds a checked-in mirror.

## File Structure

- Create: `vendor/` (entire tree, committed)
- Modify: `.gitignore` (remove any `vendor/` ignore if present — at onboard time
  `.gitignore` lists only `bin/`, `*.sanitize`, `tests/pkg/`, `proto/`, so no
  change needed; verify)
- Modify: `Makefile` (optionally pass `-mod=vendor` to `go install`/`go build`)

## Tasks

### Task 1: Run vendoring (after deps are refreshed)
- [ ] Confirm `2026-05-20-onboard-dep-refresh.md` is done — vendoring stale deps
  wastes commit churn.
- [ ] `go mod tidy`
- [ ] `go mod vendor`
- [ ] Verify: `test -d vendor && test -f vendor/modules.txt`

### Task 2: Build with vendor
- [ ] `go build -mod=vendor ./...` exits 0
- [ ] `go test -mod=vendor ./... -count=1` exits 0 (or no-tests if scaffolding plan not yet done)

### Task 3: Confirm `.gitignore` does not exclude vendor
- [ ] `grep -E '^vendor/?$' .gitignore` returns nothing.

### Task 4: Commit
- [ ] `git add vendor/ go.mod go.sum && git commit -m "deps: vendor dependencies"`

## Verification (end-to-end)

- [ ] Re-run `2026-05-20-onboard-precheck.md` Task 2.
- [ ] `go build -mod=vendor ./...` works without network.
- [ ] CI workflow (see `2026-05-20-onboard-ci-pipeline.md`) uses `-mod=vendor`.

## Cross-references

- Plan: `2026-05-20-onboard-dep-refresh.md` (run first).
- Plan: `2026-05-20-onboard-ci-pipeline.md` (consumes `-mod=vendor`).
