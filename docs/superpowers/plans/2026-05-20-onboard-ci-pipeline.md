# CI pipeline (GitHub Actions)

## Goal

Bootstrap a working GitHub Actions workflow with stages: lint → test → vuln →
build. Per user choice, GitHub Actions overrides default GitLab CI for this
repo.

## Context

No CI exists at onboard time. The iagen-dev `gitlab-cicd-go` skill ships a
GitLab CI template; the structure (lint+test+vuln+build, govulncheck reachable
gate, vendored build) translates directly to GitHub Actions. We adapt manually.

Job order assumption (locked by dependencies): `dep-refresh` and
`go-version-bump` plans complete first so `vendor/` exists and `go.mod` toolchain
directive matches the workflow's pinned Go version.

## File Structure

- Create: `.github/workflows/ci.yml`

## Tasks

### Task 1: Write workflow file

- [ ] Create `.github/workflows/ci.yml` with content:

```yaml
name: ci

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  ci:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: stable
          cache: true

      - name: gofmt
        run: |
          fmt_out=$(gofmt -l .)
          if [ -n "$fmt_out" ]; then
            echo "::error::gofmt diff:"
            echo "$fmt_out"
            exit 1
          fi

      - name: golangci-lint
        uses: golangci/golangci-lint-action@v6
        with:
          version: latest
          args: --timeout=5m

      - name: govulncheck
        run: |
          go install golang.org/x/vuln/cmd/govulncheck@latest
          govulncheck ./...

      - name: go test
        run: go test -mod=vendor ./... -count=1 -race

      - name: go build
        run: go build -mod=vendor ./...
```

### Task 2: Local dry-run

- [ ] `gofmt -l .` produces no output.
- [ ] `golangci-lint run ./...` exits 0.
- [ ] `govulncheck ./...` reachable count = 0.
- [ ] `go test -mod=vendor ./... -count=1 -race` exits 0 (or "no tests found" if
      testing plan still pending).
- [ ] `go build -mod=vendor ./...` exits 0.

### Task 3: Push and verify on GitHub

- [ ] Push branch and confirm workflow runs green on first execution.
- [ ] Open a trivial PR to verify the `pull_request` trigger.

### Task 4: Commit

- [ ] `git add .github/workflows/ci.yml && git commit -m "ci: add GitHub Actions pipeline"`

## Verification (end-to-end)

- [ ] Workflow file exists at `.github/workflows/ci.yml`.
- [ ] First push triggers a green run on default branch.
- [ ] Subsequent PRs run all four steps and block on failure.

## Cross-references

- Skill template (adapted): `gitlab-cicd-go`.
- Plan: `2026-05-20-onboard-linting.md`, `2026-05-20-onboard-vuln-scanning.md`,
  `2026-05-20-onboard-testing.md`, `2026-05-20-onboard-vendoring.md` (workflow
  assumes all four are in place).
