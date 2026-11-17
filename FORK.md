# kzu/xurl fork

This is a fork of [xdevplatform/xurl](https://github.com/xdevplatform/xurl) that
returns the full content of posts on every read endpoint (`read`, `search`,
`posts`, `timeline`, `mentions`, `bookmarks`, `likes`): the long-form text of
`note_tweet`, X articles (`article` with `title` and `plain_text`), and the
media of both.

Install it with:

```bash
go install github.com/kzu/xurl@latest
```

## What the fork changes

`main` is upstream's `main` plus the fork's commits on top:

- the module path and branding (`github.com/kzu/xurl`, `by @kzu`);
- the extra `tweet.fields`, `expansions` and `media.fields` of the read
  shortcuts in `api/shortcuts.go`, with their tests in `api/shortcuts_test.go`;
- this file.

Keep the fork's commits few and small: every one of them is rebased on each sync.

## Branches and tags

- `main` is rebased onto upstream on every sync and force-pushed. Do not pin
  `main` by commit: its commits are replaced on each sync.
- A release is an annotated tag `vX.Y.Z` named after the upstream version it is
  based on, such as `v1.3.4` for the fork of upstream v1.3.4. Go resolves
  `@latest` and `@vX.Y.Z` from these tags.
- Before a sync force-pushes `main`, the previous release keeps a branch
  `release/vX.Y.Z` at its commit, so a ref pinned to it keeps resolving.
  `release/v1.3.1` is the first; `backup/pre-upstream-rebase` holds the older
  `v0.1.x` line.
- Only the fork's own tags are pushed to `origin`. Upstream's tags are not
  fetched (see the setup below), and `git push --tags` is never used.

## Setup

```bash
git remote add upstream https://github.com/xdevplatform/xurl
# Upstream's tags have the same names as ours and point to other commits.
git config remote.upstream.tagOpt --no-tags
```

## Sync with upstream

```bash
git fetch upstream
git switch main
git rebase upstream/main
# Upstream code that imports github.com/xdevplatform/xurl must use the fork's path:
git grep -n "github.com/xdevplatform/xurl/" -- "*.go"
go build ./... && go vet ./... && go test ./...
```

`store` has upstream tests of POSIX file permissions that fail on Windows; run
the tests on Linux or macOS, or ignore those two on Windows.

Check against the live API that a long-form post and an article come back whole
(`--auth app` uses the bearer token):

```bash
go run . search 'url:"x.com/i/article" -is:retweet' -n 10 --auth app   # data[].article.plain_text
go run . search 'AI lang:en -is:retweet' -n 100 --auth app             # data[].note_tweet.text
go run . read <post id> --auth app
```

## Release

The workflows of `.github/workflows` publish upstream's Homebrew cask and npm
package; they are not used by the fork. A fork release is a tag and a GitHub
release:

```bash
# 1. Keep the current release reachable before main is replaced.
git branch release/v<current> v<current>
git push origin release/v<current>

# 2. Publish the rebased main.
git push --force-with-lease origin main

# 3. Tag it with the upstream version it is based on.
git tag -a v<upstream> -m "v<upstream> (kzu fork)"
git push origin refs/tags/v<upstream>
gh release create v<upstream> -R kzu/xurl --generate-notes
```

When a sync brings no new upstream release, there is no new version to tag:
keep `main` updated and release with the next upstream version.
