# Latch

Latch on to part of a log stream.

## Usage

Pipe logs through latch and pass it a regex with one capture group. Once the group matches it will follow that substring.

```
latch <regex>
```

```
tail -f -n 0 /some/logs.log | latch "(LT=.*)[^,\]]"
```

## Aliases

Some nice zsh aliases

```
tails='tail -F -n 0'

# latch onto a tag like "LT"
function latchtag {
    latch "($1=.*)[^,\]]"
}
```
