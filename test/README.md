# Expectations vs. Outcomes

Note: depth is being run in the root directory.

## Current Behaviour

```bash
$ depth ./test
./test
  ├ fmt
  ├ os
  └ github.com/tidwall/gjson
    ├ strconv
    ├ strings
    ├ time
    ├ unicode/utf16
    ├ unicode/utf8
    ├ unsafe
    ├ github.com/tidwall/match
      └ unicode/utf8
    └ github.com/tidwall/pretty
      ├ bytes
      ├ encoding/json
      ├ sort
      └ strconv
14 dependencies (11 internal, 3 external, 0 testing).
```

## Expected Behaviour | Stage 1 ✅

```bash
$ depth ./test
./test
  └ github.com/tidwall/gjson
    ├ github.com/tidwall/match
    └ github.com/tidwall/pretty
3 dependencies (0 internal, 3 external, 0 testing).
```

## Expected Behaviour | Stage 2

```bash
$ depth ./test
./test
  └ github.com/tidwall/gjson vX.X.X
    ├ github.com/tidwall/match vX.X.X
    └ github.com/tidwall/pretty vX.X.X
3 dependencies (0 internal, 3 external, 0 testing).
```
