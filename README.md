# 🧢 blind-guardian

[![codecov](https://codecov.io/gh/eser/blind-guardian/branch/main/graph/badge.svg?token=w6s3ODtULz)](https://codecov.io/gh/eser/blind-guardian)
[![Build Pipeline](https://github.com/eser/blind-guardian/actions/workflows/build.yml/badge.svg)](https://github.com/eser/blind-guardian/actions/workflows/build.yml)

`blind-guardian` project maintains safety lists for now, its scope can be changed in near future.

## 🙋🏻 FAQ

### Want to report a bug or request a feature?

If you're going to report a bug or request a new feature, please ensure first
that you comply with the conditions found under
[@eser/directives](https://github.com/eser/blind-guardian/blob/dev/pkg/directives/README.md).
After that, you can report an issue or request using
[GitHub Issues](https://github.com/eser/blind-guardian/issues). Thanks in advance.

### Want to contribute?

It is publicly open for any contribution from the community. Bug fixes, new
features and additional components are welcome.

If you're interested in becoming a contributor and enhancing the ecosystem,
please start by reading through our [CONTRIBUTING.md](./.github/CONTRIBUTING.md).

If you're not sure where to begin, take a look at the
[issues](https://github.com/eser/blind-guardian/issues) labeled `good first issue` and
`help wanted`. Reviewing closed issues can also give you a sense of the types of
contributions we're looking for and you can tackle.

If you're already an experienced OSS contributor, let's take you to the shortest
path: To contribute to the codebase, just fork the repo, push your changes to
your fork, and then submit a pull request.

### Requirements

- Golang 1.23 or higher (https://go.dev/)

### Usage

To get started with blind-guardian, follow these steps:

1. Initialize the project:
```bash
make init
```

2. Update the disposable email domain blocklist:
```bash
bg update           # Normal update
bg update --force   # Force update regardless of last update time
```

3. Compare differences between the last two updates:
```bash
bg diff
```

### Versioning

This project follows [Semantic Versioning](https://semver.org/). For the
versions available, see the
[tags on this repository](https://github.com/eser/blind-guardian/tags).

### License

This project is licensed under the Apache 2.0 License. For further details,
please see the [LICENSE](LICENSE) file.

### To support the project...

[Visit my GitHub Sponsors profile at github.com/sponsors/eser](https://github.com/sponsors/eser)
