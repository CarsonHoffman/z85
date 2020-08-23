# z85

z85 is a dead-simple tool to encode into and decode from the [Z85 format](http://rfc.zeromq.org/spec:32/Z85) often seen in ZeroMQ. It features a single optional command for the mode (encoding is assumed if not present), reading data from stdin and writing to stdout.

## Installation

If you have Go installed, simply run:

```sh
go get github.com/CarsonHoffman/z85
```

If you'd rather install a binary, head over to the [releases page](https://github.com/CarsonHoffman/z85/releases) and drop the appropriate binary for your platform somewhere included in your PATH.

## Usage

### Encoding

Using z85 couldn't be simpler—to encode data, all that needs to be done is to run the `z85` command, supplying data via stdin. Here's an example of encoding some public key data into Z85:

```sh
$ z85 < pubkey
@HtIbum3ML:*am5qH9Dvp3dPsROdqfC5-4pkIDRU
```

Take care that the length of your data is a multiple of 4 as required by the Z85 format; tools like `echo` may unexpextedly append a newline to your input (for `echo`, supply the `-n` flag to disable this behavior). Here's an example of using `echo` to pipe in data:

```sh
$ echo -n "testdata" | z85
By/JnwmoN*
```

### Decoding

Decoding data is *almost* as simple as encoding; we simply need to supply an extra argument to specify that we're decoding: `z85 decode` (or `z85 d` for short). Here are a few examples of the inverses of above:

```sh
$ echo -n "@HtIbum3ML:*am5qH9Dvp3dPsROdqfC5-4pkIDRU" | z85 d > pubkey_decoded
$ diff pubkey pubkey_decoded

```

```sh
$ echo -n "By/JnwmoN*" | z85 d
testdata
```

---

And that's all you need to know! Get out there and make some awesome stuff in ZeroMQ; if you ever need to get the Z85 encoding of a public key, `z85` will be waiting.

Thanks to [tilinna](https://github.com/tilinna) for their [package](https://github.com/tilinna/z85) which provides the actual Z85 functionality for this tool.
