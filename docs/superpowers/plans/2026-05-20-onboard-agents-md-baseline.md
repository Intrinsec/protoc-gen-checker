# AGENTS.md baseline

## Goal

Write `AGENTS.md` declaring tier/type, mandatory sections, and carve-outs so future
iagen-dev skill runs (lint, vuln, update-project) have an explicit contract to
diff against.

## Context

Project had no `AGENTS.md` before onboarding (2026-05-20). Tier B Go CLI requires
mandatory baseline sections: Language, Workflow Skills, Linting, Vuln scanning,
Testing, Vendoring, CI pipeline, Dev guide. All adopted (no carve-outs).

## File Structure

- Create: `AGENTS.md` (repo root)

## Tasks

### Task 1: Confirm AGENTS.md exists and is current
- [x] File written by `/isec-iagen_dev-onboard-project` on 2026-05-20
- [ ] Verify content:
  - `grep -c '^## ' AGENTS.md` returns ≥ 10 (one per mandatory section + Carve-outs + Plan index)
  - `grep -F 'Tier:** B' AGENTS.md` returns 1
  - `grep -F 'Type:** 2' AGENTS.md` returns 1

### Task 2: Lock the file in git
- [ ] `git add AGENTS.md && git commit -m "docs: add AGENTS.md (tier B Go CLI baseline)"`

## Verification (end-to-end)

- [ ] `AGENTS.md` present at repo root.
- [ ] Plan index in AGENTS.md lists this and every other plan emitted by onboard.
- [ ] Re-running `/isec-iagen_dev-update-project` reports "no drift" (or only
  drift introduced by intentional edits since 2026-05-20).

## Cross-references

- Standard: `dev-setup-project` template, tier B Go CLI mapping.
- Related skill: `dev-update-project` (refresh AGENTS.md against new rules).
