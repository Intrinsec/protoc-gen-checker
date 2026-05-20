# Vulnerability scanning

## Goal

Wire `govulncheck` into Makefile + CI so vulnerable dependencies surface on every
push, not only when somebody remembers to run it.

## Context

`govulncheck` is installed (v1.2.0) and baseline reachability scan returns 0
findings. 3 vulnerabilities exist in imported packages and 8 in transitive
modules — non-reachable but expected to drop further after `2026-05-20-onboard-dep-refresh.md`.

## File Structure

- Modify: `Makefile` (add `vuln` target)
- Reference (created by `2026-05-20-onboard-ci-pipeline.md`): `.github/workflows/ci.yml`

## Tasks

### Task 1: Add Makefile target
- [ ] Append to `Makefile`:
```
.PHONY: vuln
vuln:
	@govulncheck ./...
```
- [ ] Verify: `make vuln` exits 0 (reachable findings = 0 at baseline).

### Task 2: Establish policy
- [ ] Document in `AGENTS.md` Vuln scanning section: any **reachable** finding
  fails CI; non-reachable findings are tracked but do not block.
- [ ] Verify: `grep -F 'reachable' AGENTS.md` returns at least one hit.

### Task 3: Commit
- [ ] `git add Makefile AGENTS.md && git commit -m "vuln: add govulncheck make target"`

## Verification (end-to-end)

- [ ] Re-run `2026-05-20-onboard-precheck.md` Task 2 — `govulncheck ./...` reachable = 0.
- [ ] `make vuln` runs locally without error.

## Cross-references

- Skill: `govulncheck` (run scan), `govulncheck-install` (install).
- Plan: `2026-05-20-onboard-ci-pipeline.md` (adds the CI step that runs `make vuln`).
