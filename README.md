## Usage


Example Usage
```
go build

sovoip -user torsten -pw "mypassword" -terminal torsten_festnetz -target "#11"
```

### Options

| Option    | Default | Description             |
| --------- | ------- | ----------------------- |
| user      | -       | SIP Username            |
| password  | -       | SIP Password            |
| terminal  | -       | SIP Terminal            |
| server    | -       | SIP Server              |
| target    | -       | Target Tel or Terminal  |
| unregister| false   | Unregister old Terminal |

Options that have default values are not required.
