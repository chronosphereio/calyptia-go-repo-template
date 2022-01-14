# Go Repository Template

[![CI Status](https://img.shields.io/github/workflow/status/calyptia/go-repo-template/ci)](https://github.com/calyptia/go-repo-template/actions?query=workflow%3Aci+branch%3Amain)
[![Codecov](https://codecov.io/gh/calyptia/go-repo-template/branch/main/graph/badge.svg)](https://codecov.io/gh/calyptia/go-repo-template)

This is a GitHub repository template for Golang projects.

1.CI

- golangci-lint

- codecov

- codeQL

2.Release

- goreleaser

3.Dependabot

- gomod

- github actions

It also contains templates for:

1.Issues

2.Pull requests

3.Feature requests

## Usage

1.Use the infra repository [Calyptia Infra](https://github.com/calyptia/infra/blob/master/terraform/github.tf#L134),
add your new repository with the template set to use `calyptia/go-repo-template`.

2.Sign up on [Codecov](https://codecov.io/) and configure [Codecov GitHub Application](https://github.com/apps/codecov)
for all repositories.

3.Use the infra repository [Calyptia Infra](https://github.com/calyptia/infra/blob/master/terraform/github.tf#L134),
add your new repository with the template set to use `calyptia/go-repo-template`.

4.Replace all occurrences of `calyptia/go-repo-template` to `calyptia/repo_name`
on all files.

5.Update the following files:

   -[README.md](README.md)

   -[CONTRIBUTING.md](CONTRIBUTING.md)

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.
