# Changes from upstream

This repository is a frozen fork of etcd used for Sourcegraph QA pipelines. The following directive was added to the `go.mod` file in the following commits. The rest of the repository content is the same as upstream (but frozen).

```
replace (
    go.uber.org/zap => github.com/sourcegraph-testing/zap v1.14.1
)
```

`fb77f9b1d56391318823c434f586ffe371750321` -> `4397ceb9c11be0b3e9ee0111230235c868ba581d`
`1044a8b07c56f3d32a1f3fe91c8ec849a8b17b5e` -> `bc588b7a2e9af4f903396cdcf66f56190b9e254f`
`dfb0a405096af39e694a501de5b0a46962b3050e` -> `ad7848014a051dbe3fcd6a4cff2c7befdd16d5a8`
