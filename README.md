# gh-star-history

:octocat: Show star history of repositories. :star:

## Usage

``` console
$ gh star-history --owner cli --repo cli --per-month
Aggregate stars of cli/cli ...
2019-10 1
2020-01 513
2020-02 5462
2020-03 985
2020-04 584
2020-05 1285
2020-06 942
2020-07 608
2020-08 617
2020-09 6586
2020-10 1221
2020-11 702
2020-12 699
2021-01 568
2021-02 448
2021-03 521
2021-04 970
2021-05 456
2021-06 352
2021-07 361
2021-08 989
2021-09 500
2021-10 521
2021-11 410
2021-12 404
```

``` console
$ gh star-history --owner cli --per-year
Aggregate stars of cli/browser ...
Aggregate stars of cli/cli ...
Aggregate stars of cli/crypto ...
Aggregate stars of cli/gh-extension-precompile ...
Aggregate stars of cli/go-gh ...
Aggregate stars of cli/oauth ...
Aggregate stars of cli/safeexec ...
Aggregate stars of cli/scoop-gh ...
Aggregate stars of cli/shurcooL-graphql ...
Aggregate stars of cli/survey ...
2019    1
2020    20469
2021    6683
```

## Install

`gh-grep` can be installed as a standalone command or as [a GitHub CLI extension](https://cli.github.com/manual/gh_extension)

### Install as a standalone command

Run `gh-star-history` instead of `gh star-history`.

**deb:**

Use [dpkg-i-from-url](https://github.com/k1LoW/dpkg-i-from-url)

``` console
$ export GH-STAR-HISTORY_VERSION=X.X.X
$ curl -L https://git.io/dpkg-i-from-url | bash -s -- https://github.com/k1LoW/gh-star-history/releases/download/v$GH-STAR-HISTORY_VERSION/gh-star-history_$GH-STAR-HISTORY_VERSION-1_amd64.deb
```

**RPM:**

``` console
$ export GH-STAR-HISTORY_VERSION=X.X.X
$ yum install https://github.com/k1LoW/gh-star-history/releases/download/v$GH-STAR-HISTORY_VERSION/gh-star-history_$GH-STAR-HISTORY_VERSION-1_amd64.rpm
```

**apk:**

Use [apk-add-from-url](https://github.com/k1LoW/apk-add-from-url)

``` console
$ export GH-STAR-HISTORY_VERSION=X.X.X
$ curl -L https://git.io/apk-add-from-url | sh -s -- https://github.com/k1LoW/gh-star-history/releases/download/v$GH-STAR-HISTORY_VERSION/gh-grep_$GH-GREP_VERSION-1_amd64.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/gh-grep
```
