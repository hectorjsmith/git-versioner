# Git Versioner

Small CLI application to manage releases using git tags and branches.

## Instructions

The idea behind this tool is to leverage git tags as the system of record for a project's version. The idea being that version numbers are not stored in files within the repository but are instead read from the current commit. This avoids having to manually bump the version number.

The tool has two main modes, `release` and `fix` as well as an option to view the latest version.

**Latest**

This mode simply prints out the latest published version (as understood by the tool).

```
NAME:
   git-versioner.bin latest - Show latest version info

USAGE:
   git-versioner.bin latest [command options] [arguments...]

DESCRIPTION:
   Show the latest version for this repository.
   The version data is parsed from git tags found in the repository.

   By default prints the version string (e.g. 3.12.1).

OPTIONS:
   --verbose, -v  print more useful information about the latest version (default: false)
   --tag, -t      show the latest version tag instead of the parsed version info (default: false)
   --help, -h     show help (default: false)

```

**Release**

This mode will automatically create a new version tag. By default, it will attempt the read the new version from the branch name (e.g. when run on a branch named `release/v1.2.3` it will create a `v1.2.3` tag).

Alternatively, the `release` mode can be used to increment the major or minor versions (using the `--minor` or `--major` flags).

It is also possible to include a message for that version using the `--message` flag. This will create an annotated tag with that message.

The `--test` flag can be used instead of `--major` or `--minor` to create a test version tag. This tag will use `git describe` to generate a unique identifier for that version. This is to allow releasing test/preview versions without incrementing the actual version number.

```
NAME:
   git-versioner.bin release - Create new version tag

USAGE:
   git-versioner.bin release [command options] [arguments...]

DESCRIPTION:
   Create a new version git tag named by taking the latest version and incrementing it.
   This command assumes the use of semantic versioning. The version string is parsed as: <major>.<minor>.<bugfix>
   Options are available to increment the major or minor versions.
   If no options are provided, the version number to use will be parsed from the current branch.
   For example, if run on a branch named 'release/v1.2.3', the new tag would be 'v1.2.3'.

   The repository must not have un-staged changes - i.e. the repo cannot be dirty

OPTIONS:
   --message value, -m value  Message to put in git tag. Using this will create an annotated tag.
   --major                    major release (v1.5.10 -> v2.0.0) (default: false)
   --minor                    minor release (v1.5.10 -> v1.6.0) (default: false)
   --test                     test release (v1.5.10-5-g600d3f2) (default: false)
   --help, -h                 show help (default: false)
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
   git-versioner.bin fix - Create fix branch

USAGE:
   git-versioner.bin fix [command options] [arguments...]

DESCRIPTION:
   Create a fix branch for the specified version (or latest version).
   This command will checkout the selected version (based on the corresponding git tag) and create a new fix branch.

   The repository must not have un-staged changes - i.e. the repo cannot be dirty

OPTIONS:
   --version value, -v value  version to fix (e.g. '1.2.0')
   --help, -h                 show help (default: false)
```
