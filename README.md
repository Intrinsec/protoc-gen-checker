# Protoc Gen Checker

Protoc plugin to execute checks and rules.

This project uses [protoc-gen-star](https://github.com/lyft/protoc-gen-star) to ease code generation.

## How to use it

See `./tests/incident.proto` for example on how to use it.

## Options

The plugin supports the following options.

Entities having options to disable the validation MUST have a reason explaining why. This reason MUST be given in a leading comments starting with `// No Validation Reason: `.

### At file level

`disable_file_validate`: indicates the plugin should not check if validation is correctly setup on the file.

### At message level

No specific option are defined since `protoc-gen-validate` already has option to disable/ignore the validation. The plugin check these options to know what to do.

### At field level

`disable_field_validate`: indicates the validation is disabled on the field.

## Tests

Do a `make test` to run the plugin on some examples proto files. This test is supposed to fail because it illustrate various wrong ways of using the plugin.
