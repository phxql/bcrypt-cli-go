# BCrypt CLI

Usage: `bcrypt-cli-go [password] [rounds]`

* `password` defaults to `-`
* `rounds` defaults to `12`.

If `password` is `-`, read the password from stdin.

## Examples

```
bcrypt-cli-go
```

Reads the password from stdin, uses 12 rounds.

```
bcrypt-cli-go foobar
```

Uses `foobar` as the password with 12 rounds.

```
bcrypt-cli-go foobar 14
```

Uses `foobar` as the password with 14 rounds.

```
bcrypt-cli-go - 14
```

Reads the password from stdin, uses 14 rounds.

## License

[GPLv3](https://www.gnu.org/licenses/gpl-3.0.en.html)