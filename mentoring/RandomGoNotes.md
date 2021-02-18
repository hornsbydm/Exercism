Go Notes
========

Delve
-----

### Setting Break Point

#### On Function
``` bash
$ dlv test .
Type 'help' for list of commands.
(dlv) break luhn.Valid
(dlv) continue
...
```

#### In Code
``` go
runtime.Breakpoint()
```

#### On line in file
``` bash
(dlv) break <breakpoint_name> <filename_pattern>:<line_number>
```