# Developer guide

## Goal

Write `docs/DEVELOPMENT.md` so a new contributor can build, regenerate
protobuf, run tests, cut a release, and recognise the project's two surprising
bits (the deliberately-failing `make test`, and the curl-fetched
`proto/validate/` directory).

## Context

Tier B Go projects require `docs/DEVELOPMENT.md`. Onboarding has README.md
(thin) and Makefile (working but undocumented). The dev guide consolidates the
"how to work on this repo" knowledge.

## File Structure

- Create: `docs/DEVELOPMENT.md`

## Tasks

### Task 1: Draft DEVELOPMENT.md

- [ ] Write `docs/DEVELOPMENT.md` covering at minimum these sections:

  - **Build**: `make build`, output `bin/protoc-gen-checker`.
  - **Install**: `make install` (drops binary into `$GOBIN`; protoc finds it
    via PATH).
  - **Regenerate `checker/checker.pb.go`**: explain dependency chain
    (`bin/protoc-gen-go` → `proto/validate/` fetched via `make proto/validate`).
  - **Run fixture test**: `make test`. Highlight that this command **is
    expected to print failures** — they are the demonstrations of misuse the
    plugin reports. Quote the banner so a contributor doesn't panic.
  - **Run unit tests**: `go test ./... -count=1 -race` (once
    `2026-05-20-onboard-testing.md` is applied).
  - **Lint**: `make lint`.
  - **Vuln scan**: `make vuln`.
  - **Vendor refresh**: `go mod tidy && go mod vendor && git add vendor go.mod go.sum`.
  - **Release flow**: tag (`vX.Y.Z`), push tag, CI builds binary artifacts.
    (Detail target: signed releases not yet adopted — link to a future plan
    when adopted.)
  - **Common pitfalls**:
    - `proto/validate/` is **fetched from upstream `protoc-gen-validate`** via
      `make proto/validate`. It is `.gitignore`d. Run that target once after
      clone.
    - `make distclean` wipes generated proto and binaries; useful when bumping
      `PROTOC_GEN_VALIDATE` version in the Makefile.
    - `checker/checker.pb.go` is generated — never hand-edit.
- [ ] Verify: `wc -l docs/DEVELOPMENT.md` ≥ 60.

### Task 2: Cross-link from README

- [ ] Add a "Development" pointer near the top of `README.md`:
  `See [docs/DEVELOPMENT.md](docs/DEVELOPMENT.md) for build, test, and release instructions.`
- [ ] Verify: `grep -F 'docs/DEVELOPMENT.md' README.md` returns 1.

### Task 3: Commit

- [ ] `git add docs/DEVELOPMENT.md README.md && git commit -m "docs: add DEVELOPMENT.md"`

## Verification (end-to-end)

- [ ] `docs/DEVELOPMENT.md` present, ≥ 60 lines.
- [ ] Following the guide on a fresh clone produces a working binary and
  reproduces the `make test` expected-failure output.

## Cross-references

- Standard: `dev-setup-project` tier B dev-guide template (read for the
  canonical section order).
