# rnm

Batch rename files and folders.
This project is a Go implementation of [75lb/renamer](https://github.com/75lb/renamer).

## Example

```sh
$ rnm --find 'bad' --replace 'good' *
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td>
                <pre><code>.
├── A poem-bad.txt
├── A story-bad.txt
                </code></pre>
            </td>
            <td>
                <pre><code>.
├── A poem-good.txt
├── A story-good.txt
                </code></pre>
            </td>
        </tr>
    </tbody>
</table>

## Installation

```sh
$ go get github.com/ryym/rnm/cmd/rnm
```
