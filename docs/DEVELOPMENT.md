# Development guide — protoc-gen-checker

Audience: contributors. Pair with `AGENTS.md` for the project rules.

## Toolchain

- **Go:** latest stable (>= 1.26).
- **protoc:** any recent release that supports the protobuf descriptor APIs used
  by `protoc-gen-validate` v1.3.x and `lyft/protoc-gen-star/v2`.
- **golangci-lint:** any v2.x. `make lint` uses the system binary.
- **govulncheck:** any v1.x. `make vuln` uses the system binary.

## One-time setup after clone

```sh
make proto/validate   # fetches validate.proto from the protoc-gen-validate release tarball
```

`proto/validate/` is `.gitignore`d on purpose — the validate proto definitions
come from an external release tagged by `PROTOC_GEN_VALIDATE` in the `Makefile`
(currently `1.3.3`). Bump that variable when you want a newer release.

## Build

```sh
make build       # produces bin/protoc-gen-checker
make install     # go install . into $GOBIN
```

`protoc` finds the plugin via `$PATH`, so `make install` is enough for a
host-wide invocation. For local invocation pass `--plugin=protoc-gen-checker=./bin/protoc-gen-checker`.

## Regenerate the checker proto

`checker/checker.pb.go` is generated from `checker/checker.proto`. Do **not**
hand-edit it. To regenerate:

```sh
make bin/protoc-gen-go        # fetched once
make checker/checker.pb.go
```

The first target installs `google.golang.org/protobuf/cmd/protoc-gen-go` into
the local `bin/`. The second invokes `protoc` with the import-path remappings
declared in `GO_IMPORT_SPACES`.

## Run unit tests

```sh
make unit        # go test ./... -count=1 -race
```

These cover pure-Go helpers in `checker.go` (`getNoValidationReason`). Visitor
methods are exercised by the integration fixture below.

## Run fixture integration test

```sh
make test
```

**This command is expected to print failures.** It runs the plugin against
`tests/*.proto` in strict mode; those fixtures intentionally illustrate every
misuse the plugin reports. The banner repeats the warning:

```
/////////////////////////////////////////////////////////////////////////////////////////////
This test is supposed to FAIL. It illustrate the various good/wrong ways of using this plugin.
/////////////////////////////////////////////////////////////////////////////////////////////
```

A real regression looks like: the failure list changing, or the plugin crashing
instead of reporting findings. Use `git diff` against a known-good run to spot
real changes.

## Lint

```sh
make lint        # golangci-lint run ./...
```

Config is `.golangci.yml` (Standard preset). `ireturn` and `nilnil` are disabled
because the PGS visitor API contractually returns `pgs.Visitor` and signals
"stop descending" with `nil, nil`.

## Vulnerability scan

```sh
make vuln        # govulncheck ./...
```

Reachable findings (`Your code is affected by N vulnerabilities`) **must** be
zero. Non-reachable findings in imported packages or transitive modules are
tracked but do not block — they get cleared by dep refreshes.

Reachable Go stdlib findings clear after upgrading the system `go` toolchain.

## Release flow

1. Bump dep versions if needed (`go get …@latest`; `go mod tidy`).
2. `make lint && make vuln && make unit && make test` — last one expected to
   FAIL with the standard banner.
3. Tag: `git tag v<MAJOR>.<MINOR>.<PATCH> && git push --tags`.
4. CI builds and publishes the binary artifact for the tag.

Signed releases / SBOM publication are not adopted yet (tier B). Add a plan
under `docs/superpowers/plans/` when they become a requirement.

## Common pitfalls

- **`proto/validate/` missing after fresh clone.** Run `make proto/validate`.
- **`make distclean` wipes everything generated**, including `checker/checker.pb.go`
  and the fetched validate proto. Useful when bumping `PROTOC_GEN_VALIDATE` in
  the `Makefile`.
- **Editing `checker/checker.pb.go`.** Never. It is generated.
- **Lint complains about generated code.** `.golangci.yml` excludes `*.pb.go`,
  `vendor/`, and `third_party/`. If a new generator is added, append its
  output path regex to both `linters.exclusions.paths` and
  `formatters.exclusions.paths`.
- **Vendor mode** is carved out (see `AGENTS.md`). Do not run `go mod vendor`
  unless reinstating the carve-out — `vendor/` is in `.gitignore`.
