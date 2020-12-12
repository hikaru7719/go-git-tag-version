# go-git-tag-version

go-git-tag-version supports versioning of git tag.
You can do versioning using go-git-tag-version with --major,--minor,--patch options like yarn version command.

## Usage

You can use this cli in your local git repository.
If you use this cli in except git repository, cli wil fail with error.
You must install git command before you use this cli.

Command line options is below.

```
Usage:
  go-git-tag-version [flags]

Flags:
  -d, --dry-run   execute dry-run mode
  -h, --help      help for go-git-tag-version
  -a, --major     increment major version
  -b, --minor     increment minor version
  -c, --patch     increment patch version
```

## Example

Current version is v1.0.0.

```
$ git tag
v1.0.0
```

You execute cli app.

```
$ go-git-tag-version --major
2020/12/12 13:07:30 new version is v2.0.0
```

After execution, Current version is v2.0.0.

```
$ git tag
v1.0.0 v2.0.0
```
