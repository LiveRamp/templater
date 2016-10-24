# templater: the power of Go templates available on the command line

Ever wanted to use Go's [text/template](https://golang.org/pkg/text/template/)
in shell tooling instead of having to use a
[here document](http://www.tldp.org/LDP/abs/html/here-docs.html)?
Well, now you can. templater is a command line tool that accepts an arbitrary
template input file, and a JSON data file to use as the context for rendering
the template.

Now you can drive templates that use inline conditionals, loops, format strings
from the convenience of any script that allows you to fork out to templater.

*This tool is an open source project of [Arbor Technologies](https://arbor.io),
based in New York City. Want to hack on cool Go projects at adtech scale?
[Come work for us!](https://arbor.io/careers/)*

### Building

```
$ go get -u github.com/pippio/templater
$ go install github.com/pippio/templater
```

### Usage
```
Usage of templater:
  -data string
        Data file to use as template context
  -template string
        Template to render
```

### Example

Example template input file:

```
I have three students:
{{ range $index, $element := .Students }}
One student is named {{ $element }}.
{{ end }}
```

Example JSON data file:

```
{"Students": ["Curly", "Larry", "Moe"]}
```

Rendering the template:

```
$ templater -template example.tmpl -data example.json
I have three students:

One student is named Curly.

One student is named Larry.

One student is named Moe.
```

### What's Next

* Load multiple templates into the context so you can include other templates
  from the root template.
* Multiple data files with precedence.
* Additional template context functions defined in Go to reach parity with
  Erubis, Jinja, etc.
