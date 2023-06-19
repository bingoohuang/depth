# Expectations vs. Outcomes

A json version of the tree(s) below can be exported using `-json` flag.

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
$ depth -std ./test
./test
  └ github.com/tidwall/gjson
    ├ github.com/tidwall/match
    └ github.com/tidwall/pretty
3 dependencies (0 internal, 3 external, 0 testing).
```

## Expected Behaviour | Stage 2 ✅

```bash
$ depth -std -v ./test
./test
  └ github.com/tidwall/gjson@v1.14.4
    ├ github.com/tidwall/match@v1.1.1
    └ github.com/tidwall/pretty@v1.2.0
3 dependencies (0 internal, 3 external, 0 testing).
```
