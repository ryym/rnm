# rnm

[![travis][travis-badge]](https://travis-ci.org/ryym/rnm)

[travis-badge]: https://travis-ci.org/ryym/rnm.svg?branch=travis

Batch rename files and folders.

**Note that this project is a Go implementation of [75lb/renamer](https://github.com/75lb/renamer)**
and most of the README contents are copied from the original repository.

## Install

```sh
$ go get github.com/ryym/rnm/cmd/rnm
```

## Usage

```sh
$ rnm [options] <files>
```

### -f, --find &lt;string&gt;

The find string, or regular expression when `--regex` is set. If not set, the whole filename will be replaced.

### -r, --replace &lt;string&gt;

The replace string. With `--regex` set, `--replace` can reference parenthesised substrings from `--find` with $1, $2, $3 etc. If omitted, defaults to a blank

### -e, --regex

When set, --find is intepreted as a regular expression.

### -d, --dry-run

Used for test runs. Set this to do everything but rename the file.

**Don't forget to test your rename first using `--dry-run`!**

## Examples

Some real-world examples.

**Windows users**: the single-quotation marks used in the example commands below are for bash (Mac/Linux) users, please replace these with double-quotation marks on Windows.

### Simple replace

```sh
$ rnm --find '[bad]' --replace '[good]' *
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── A poem [bad].txt
├── A story [bad].txt
            </code></pre></td>
            <td><pre><code>.
├── A poem [good].txt
├── A story [good].txt
            </code></pre></td>
        </tr>
    </tbody>
</table>

### Strip out unwanted text

```sh
$ rnm --find 'Season 1 - ' *
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── Season 1 - Some crappy episode.mp4
├── Season 1 - Load of bollocks.mp4
            </code></pre></td>
            <td><pre><code>.
├── Some crappy episode.mp4
├── Load of bollocks.mp4
            </code></pre></td>
        </tr>
    </tbody>
</table>

### Simple filename cleanup

```sh
$ rnm --regex --find '.*_(\d+)_.*' --replace 'Video $1.mp4' *
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [ag]_Annoying_filename_-_3_[38881CD1].mp4
├── [ag]_Annoying_filename_-_34_[38881CD1].mp4
├── [ag]_Annoying_filename_-_53_[38881CD1].mp4
            </code></pre></td>
            <td><pre><code>.
├── Video 3.mp4
├── Video 34.mp4
├── Video 53.mp4
            </code></pre></td>
        </tr>
    </tbody>
</table>

### if not already done, add your name to a load of files
```sh
$ rnm --regex --find '(data\d)(\.\w+)' --replace '$1 (checked by Lloyd)$2' *
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── data1.csv
├── data2 (checked by Lloyd).csv
├── data3.xls
            </code></pre></td>
            <td><pre><code>.
├── data1 (checked by Lloyd).csv
├── data2 (checked by Lloyd).csv
├── data3 (checked by Lloyd).xls
            </code></pre></td>
        </tr>
    </tbody>
</table>

### rename files and folders, recursively

```sh
$ rnm --find 'pic' --replace 'photo' '**/*'
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── pic1.jpg
├── pic2.jpg
└── pics
    ├── pic3.jpg
    └── pic4.jpg
            </code></pre></td>
            <td><pre><code>.
├── photo1.jpg
├── photo2.jpg
└── photos
    ├── photo3.jpg
    └── photo4.jpg
            </code></pre></td>
        </tr>
    </tbody>
</table>

### prefix files and folders, recursively

```sh
$ rnm --regex --find '^' --replace 'good-' '**/*'
```

<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── pic1.jpg
├── pic2.jpg
└── pics
    ├── pic3.jpg
    └── pic4.jpg
            </code></pre></td>
            <td><pre><code>.
├── good-pic1.jpg
├── good-pic2.jpg
└── good-pics
    ├── good-pic3.jpg
    └── good-pic4.jpg
            </code></pre></td>
        </tr>
    </tbody>
</table>
