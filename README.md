# Git Versioner

Small CLI application to manage releases using git tags and branches.

## Instructions

The idea behind this tool is to leverage git tags as the system of record for a project's version. The idea being that version numbers are not stored in files within the repository but are instead read from the current commit. This avoids having to manually bump the version number.

The tool has two main modes, `release` and `fix` as well as an option to view the latest version.

**Latest**

This mode simply prints out the latest published version (as understood by the tool).

```
NAME:
   git-versioner.bin latest - return the latest version info (highest value) - this is the version that will be incremented for release

USAGE:
   git-versioner.bin latest [command options] [arguments...]

OPTIONS:
   --verbose, -v  print more useful information about the latest version
   --tag, -t      only show the latest tag, not the version info
```

**Release**

This mode will automatically create a new version tag. By default it will attempt the read the new version from the branch name (e.g. when run on a branch named `release/v1.2.3` it will create a `v1.2.3` tag).

Alternatively, the `release` mode can be used to increment the major or minor versions (using the `--minor` or `--major` flags).

It is also possible to include a message for that version using the `--message` flag. This will create an annotated tag with that message.

The `--test` flag can be used instead of `--major` or `--minor` to create a test version tag. This tag will use `git describe` to generate a unique identifier for that version. This is to allow releasing test/preview versions without incrementing the actual version number.

```
NAME:
   git-versioner.bin release - release a new version - by default the branch name is used to parse the new version

USAGE:
   git-versioner.bin release [command options] [arguments...]

OPTIONS:
   --message value  Message to put in git tag. Using this will create an annotated tag.
   --major          major release (v1.5.10 -> v2.0.0)
   --minor          minor release (v1.5.10 -> v1.6.0)
   --test           test release (v1.5.10-5-g600d3f2)
```

**Fix**

This mode is to be used when going back to fix a previous version.
This mode will checkout the tag for the version to fix (by default the latest version) and create a new fix branch for it.

For example, to fix `v1.2.0` the tool will:
1. Checkout the `v1.2.0` tag
2. Create a new branch with the new version name: `rel/v1.2.1`

Once the fix has been done, the tool can be used again in `release` mode to create the new version tag by reading the branch name.

*Note*: The `fix` command will ignore any test version tags (see above). It will only create fix branches for proper version tags.

```
NAME:
   git-versioner.bin fix - create a fix branch for an existing version

USAGE:
   git-versioner.bin fix [command options] [arguments...]

OPTIONS:
   --version value, -v value  version to fix (e.g. '1.2.0')
```
