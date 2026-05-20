# Dependency refresh

## Goal

Bump direct and transitive dependencies to current stable versions. `govulncheck`
flags 3 imported-package + 8 module-level vulnerabilities at onboard time — none
reachable, but they should not linger.

## Context

Direct deps at onboard time (2026-05-20):

| Module | Current | Notes |
|--------|---------|-------|
| `github.com/envoyproxy/protoc-gen-validate` | `v0.6.1` | Renamed → `buf.build/go/protovalidate`; check migration path |
| `github.com/golang/protobuf` | `v1.5.2` | Deprecated → `google.golang.org/protobuf` (already required) |
| `github.com/lyft/protoc-gen-star` | `v0.6.0` | Repo moved to `github.com/lyft/protoc-gen-star/v2` |
| `google.golang.org/protobuf` | `v1.27.1` | Bump to latest stable |

## File Structure

- Modify: `go.mod`, `go.sum`
- Modify: `main.go`, `checker.go` if import paths change (e.g. lyft → v2)
- Modify: `Makefile` if PGV migration changes proto-fetch URL

## Tasks

### Task 1: Bump trivial deps
- [ ] `go get google.golang.org/protobuf@latest`
- [ ] `go get github.com/lyft/protoc-gen-star/v2@latest` (if migration is straightforward)
- [ ] `go mod tidy`
- [ ] `go build ./...` exits 0 — if it does not, see Task 2.

### Task 2: Investigate `lyft/protoc-gen-star` → v2 migration
- [ ] Check upstream release notes: https://github.com/lyft/protoc-gen-star/releases
- [ ] If v2 is the live line: rewrite imports in `main.go`, `checker.go` from
  `github.com/lyft/protoc-gen-star` → `github.com/lyft/protoc-gen-star/v2`
  (and `lang/go` subpackage path likewise)
- [ ] `go build ./...` exits 0

### Task 3: Investigate `protoc-gen-validate` migration
- [ ] Confirm whether `v0.6.x` still works with current protobuf. If yes, bump to
  latest `v0.x` patch and leave migration to a future plan.
- [ ] If `v0.6.1` no longer builds against the bumped protobuf, follow
  https://github.com/bufbuild/protovalidate-go migration guide and update
  `Makefile`'s `PROTOC_GEN_VALIDATE` URL + import paths in `checker.go` (the
  `validate.E_Disabled`, `validate.E_Ignored`, `validate.E_Rules` extensions).

### Task 4: Drop `github.com/golang/protobuf` (deprecated)
- [ ] Replace any direct uses with `google.golang.org/protobuf` equivalents.
  At onboard time only `Makefile` references `golang/protobuf/ptypes/*` via
  `GO_IMPORT_SPACES` — verify whether protoc-gen-go still needs these
  remappings or if they can be deleted.
- [ ] If no longer required, simplify `Makefile`.

### Task 5: Verify clean state
- [ ] `go mod tidy` is idempotent
- [ ] `go build ./...` exits 0
- [ ] `make test` still demonstrates the expected fixture failures (its job)
- [ ] `govulncheck ./...` shows reachable count unchanged (0) and indirect counts
  reduced

### Task 6: Commit
- [ ] `git add go.mod go.sum main.go checker.go Makefile && git commit -m "deps: refresh stale dependencies"`

## Verification (end-to-end)

- [ ] Re-run `2026-05-20-onboard-precheck.md` Task 2.
- [ ] `govulncheck ./...` "vulnerabilities in modules you require" delta ≤ 0 vs baseline.
- [ ] Plugin still passes its own `make test` fixture (which intentionally fails;
  the failure message must be unchanged in structure).

## Cross-references

- Skill: `govulncheck` (run reachable-vuln scan after each bump).
- Plan: `2026-05-20-onboard-go-version-bump.md` (must run first).
