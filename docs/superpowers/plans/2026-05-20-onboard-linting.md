# Linting

## Goal

Create `.golangci.yml` Standard preset for tier B; wire `make lint` target; gate CI.

## Context

`golangci-lint` is installed (v2.10.1 at onboard time) but no project config
exists. Baseline scan with default linters shows 0 issues, which is good
but unconstrained — a richer Standard preset is wanted so style/security/perf
checks run too.

## File Structure

- Create: `.golangci.yml` (repo root)
- Modify: `Makefile` (add `lint` target)

## Tasks

### Task 1: Generate `.golangci.yml` via skill
- [ ] Invoke `/isec-iagen_lint-go-config` and pick **Standard** preset (tier B).
- [ ] Verify file exists: `test -f .golangci.yml`
- [ ] Verify it parses: `golangci-lint config verify`

### Task 2: Run lint, fix any new findings
- [ ] `golangci-lint run ./...`
- [ ] If exit code ≠ 0: triage each finding. Fix code where straightforward, or
      add scoped `//nolint:<linter> // <reason>` (reason mandatory).
- [ ] Re-run until exit code = 0.

### Task 3: Add Makefile target
- [ ] Append to `Makefile`:
```
.PHONY: lint
lint:
	@golangci-lint run ./...
```
- [ ] Verify: `make lint` exits 0.

### Task 4: Commit
- [ ] `git add .golangci.yml Makefile <any-fixed-go-files> && git commit -m "lint: add golangci-lint Standard preset"`

## Verification (end-to-end)

- [ ] Re-run `2026-05-20-onboard-precheck.md` Task 2 — `golangci-lint run ./...` exit 0.
- [ ] `.golangci.yml` present, valid, and CI step in `2026-05-20-onboard-ci-pipeline.md` runs `make lint`.

## Cross-references

- Skill: `lint-go-config` (creates config), `lint-go` (runs lint).
