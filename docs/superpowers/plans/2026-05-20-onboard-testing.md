# Testing — scaffold first unit tests

## Goal

Bring repo from 0 `*_test.go` files to a minimal but real Go unit test suite
covering pure functions in `checker.go`. Fixture-based `make test` stays as an
integration sanity check; it does not satisfy tier B's unit test requirement.

## Context

`checker.go` contains pure logic worth direct unit tests:

- `getNoValidationReason(comments string) string` — string parser, trivial to
  test end-to-end with table tests.
- `(*checkVisitor).isValidationDisabled(...)` — branching on `ok`, `err`,
  `disabled`. Side-effects (`v.Logf`, `v.missingValidation`) can be exercised via
  a fake `pgs.DebuggerCommon`.

Visitor methods (`VisitFile`, `VisitMessage`, `VisitField`) depend on `pgs.File`/
`pgs.Message`/`pgs.Field` interfaces from `lyft/protoc-gen-star`. These are
harder to fake. Defer integration-style coverage of those to a follow-up plan
once dep refresh stabilises the PGS major version.

## File Structure

- Create: `checker_test.go` (package `main`, alongside `checker.go`)
- Reference: `checker.go:186` (`noValidationMarker`), `checker.go:189` (`getNoValidationReason`).

## Tasks

### Task 1: TDD a `getNoValidationReason` test
- [ ] Following `superpowers:test-driven-development`, create `checker_test.go`
  with table tests:
  - empty string → `""`
  - single line `// No Validation Reason: foo` → `"foo"` (note: marker
    `noValidationMarker = " No Validation Reason: "` — leading space — so the
    function expects the comment with the leading space included after `//`
    has been stripped by the caller; test inputs reflect that)
  - multi-line with marker on second line → that line's reason
  - line with marker but empty reason → `""`
  - multi-line, marker missing entirely → `""`
- [ ] Run: `go test ./... -run TestGetNoValidationReason -v` exits 0.

### Task 2: TDD `isValidationDisabled` (optional, scope cap if PGS interface is unwieldy)
- [ ] Build a minimal fake implementing only the `pgs.Entity` methods actually
  called (`File().Name()`, `SourceCodeInfo().Location().Span`, `Name()`,
  `SourceCodeInfo().LeadingComments()`) plus a fake `pgs.DebuggerCommon` that
  records `Logf` calls.
- [ ] Cover: `ok=false`, `err≠nil`, `disabled=false`, `disabled=true` + reason
  present, `disabled=true` + reason missing (must flip `missingValidation`).
- [ ] If the PGS fake balloons past ~50 lines, stop and move this task to a
  follow-up plan — table testing pure functions first is the priority.

### Task 3: Race + coverage check
- [ ] `go test ./... -count=1 -race` exits 0.
- [ ] `go test ./... -count=1 -cover` reports non-zero coverage on `checker.go`
  (specifically `getNoValidationReason`).

### Task 4: Commit
- [ ] `git add checker_test.go && git commit -m "test: add unit tests for getNoValidationReason"`

## Verification (end-to-end)

- [ ] Re-run `2026-05-20-onboard-precheck.md` Task 2 — `go test` no longer
  reports "no tests found".
- [ ] Coverage on `getNoValidationReason` ≥ all listed table-test branches.

## Cross-references

- Workflow skill: `superpowers:test-driven-development`.
- Verification skill: `superpowers:verification-before-completion`.
