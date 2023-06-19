# What's going on?

## Current Design

The `depth` parses a **go package** and spits out a _dependency tree_ of that specific package.
If the package is dependent on some external package, then it recursively parses it too.
It continues to resolve packages until only the standard libraries remain.

> Since Go doesn't support cyclic dependencies, the tree will eventually stop. [Read more](https://stackoverflow.com/questions/28256923/import-cycle-not-allowed)

`depth` also provides various flags.

- `explain {packageName}`: lists the packages that import `packageName`.
- `internal`: resolves standard library packages.
- `json`: outputs in JSON format.
- `max {count}`: resolves up to `count` height.
- `test`: resolves the testing dependencies.

## Upgrades needed

As of version `1.2.1`, the two features that are not supported are:

1. Flags to specify a `tag` or a `branch`.
2. Versioned dependency.

I aim to add these functionalities to this application.

## Approch

### Tag

- Save the working branch.
- Create a `test` branch.
- Get the `commitID` using `git show-ref {tag}`.
- Perform `git reset --hard {commitID}`.
- Run `depth .` or dependency parsing internally.
- Revert back to the working branch.

### Branch

- Save the working branch.
- Checkout to the desired branch.
- Run `depth .` or dependency parsing internally.
- Revert back to working branch.

### Dependency Version

- Run `go list -m all`

## Judging Assumptions

From an interviewer's standpoint, I am assuming this assignment was given to test abilities beyond coding. I assume they are

- The ability to be resourceful with respect to code quality and time.
- The ability to research and cherry pick features from different projects
- The ability to read, write, and work in complex and/or large code bases
