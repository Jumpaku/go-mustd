# go-mustd

[![Go Reference](https://pkg.go.dev/badge/github.com/Jumpaku/go-mustd.svg)](https://pkg.go.dev/github.com/Jumpaku/go-mustd)

go-mustd realizes easy Go scripting based on the error exit behavior like `set -e` in Shell Script.
go-mustd provides "must" versions of some standard APIs, which panic on error instead of returning errors.

## Installation

```bash
go get github.com/Jumpaku/go-mustd
```

## Features

- strconvmust: "must" version of standard strconv package
  - `func Atoi(s string) int`: "must" version of `strconv.Atoi`
  - `func ParseBool(str string) bool`: "must" version of `strconv.ParseBool`
  - `func ParseComplex(s string, bitSize int) complex128`: "must" version of `strconv.ParseComplex`
  - `func ParseFloat(s string, bitSize int) float64`: "must" version of `strconv.ParseFloat`
  - `func ParseInt(s string, base int, bitSize int) int64`: "must" version of `strconv.ParseInt`
  - `func ParseUint(s string, base int, bitSize int) uint64`: "must" version of `strconv.ParseUint`
  - `func QuotedPrefix(s string) string`: "must" version of `strconv.QuotedPrefix`
  - `func Unquote(s string) string`: "must" version of `strconv.Unquote`
  - `func UnquoteChar(s string, quote byte) (rune, bool, string)`: "must" version of `strconv.UnquoteChar`
- bytesmust: "must" version of standard bytes package
  - `type Buffer`: "must" version of `bytes.Buffer`
  - `type Reader`: "must" version of `bytes.Reader`
- iomust: "must" version of standard io package
  - `func Copy(dst Writer, src Reader) int64`: "must" version of `io.Copy`
  - `func CopyBuffer(dst Writer, src Reader, buf []byte) int64`: "must" version of `io.CopyBuffer`
  - `func CopyN(dst Writer, src Reader, n int64) int64`: "must" version of `io.CopyN`
  - `func Pipe() (*PipeReader, *PipeWriter)`: "must" version of `io.Pipe`
  - `func ReadAll(r Reader) []byte`: "must" version of `io.ReadAll`
  - `func ReadAtLeast(r Reader, buf []byte, min int) int`: "must" version of `io.ReadAtLeast`
  - `func ReadFull(r Reader, buf []byte) int`: "must" version of `io.ReadFull`
  - `func WriteString(w Writer, s string) int`: "must" version of `io.WriteString`
  - `type Closer`: "must" version of `io.Closer`
  - `type Reader`: "must" version of `io.Reader`
  - `type Writer`: "must" version of `io.Writer`
  - `type Seeker`: "must" version of `io.Seeker`
  - `type ReadCloser`: "must" version of `io.ReadCloser`
  - `type ReadSeeker`: "must" version of `io.ReadSeeker`
  - `type ReadSeekCloser`: "must" version of `io.ReadSeekCloser`
  - `type WriteCloser`: "must" version of `io.WriteCloser`
  - `type WriteSeeker`: "must" version of `io.WriteSeeker`
  - `type ReadWriter`: "must" version of `io.ReadWriter`
  - `type ReadWriteCloser`: "must" version of `io.ReadWriteCloser`
  - `type ReadWriteSeeker`: "must" version of `io.ReadWriteSeeker`
  - `type PipeReader`: "must" version of `io.PipeReader`
  - `type PipeWriter`: "must" version of `io.PipeWriter`
- osmust: "must" version of standard os package
  - `func Chdir(dir string)`: "must" version of `os.Chdir`
  - `func Chmod(name string, mode os.FileMode)`: "must" version of `os.Chmod`
  - `func Chown(name string, uid, gid int)`: "must" version of `os.Chown`
  - `func Chtimes(name string, atime, mtime time.Time)`: "must" version of `os.Chtimes`
  - `func Create(name string) *File`: "must" version of `os.Create`
  - `func CreateTemp(dir, pattern string) *File`: "must" version of `os.CreateTemp`
  - `func Executable() string`: "must" version of `os.Executable`
  - `func Getgroups() []int`: "must" version of `os.Getgroups`
  - `func Getwd() string`: "must" version of `os.Getwd`
  - `func Hostname() string`: "must" version of `os.Hostname`
  - `func Lchown(name string, uid, gid int)`: "must" version of `os.Lchown`
  - `func Link(oldname, newname string)`: "must" version of `os.Link`
  - `func Lstat(name string) os.FileInfo`: "must" version of `os.Lstat`
  - `func Mkdir(name string, perm os.FileMode)`: "must" version of `os.Mkdir`
  - `func MkdirAll(path string, perm os.FileMode)`: "must" version of `os.MkdirAll`
  - `func MkdirTemp(dir, pattern string) string`: "must" version of `os.MkdirTemp`
  - `func Open(name string) *File`: "must" version of `os.Open`
  - `func OpenFile(name string, flag int, perm os.FileMode) *File`: "must" version of `os.OpenFile`
  - `func Pipe() (*File, *File)`: "must" version of `os.Pipe`
  - `func ReadFile(name string) []byte`: "must" version of `os.ReadFile`
  - `func Readlink(name string) string`: "must" version of `os.Readlink`
  - `func Remove(name string)`: "must" version of `os.Remove`
  - `func RemoveAll(path string)`: "must" version of `os.RemoveAll`
  - `func Rename(oldpath, newpath string)`: "must" version of `os.Rename`
  - `func Setenv(key, value string)`: "must" version of `os.Setenv`
  - `func Stat(name string) os.FileInfo`: "must" version of `os.Stat`
  - `func Symlink(oldname, newname string)`: "must" version of `os.Symlink`
  - `func Truncate(name string, size int64)`: "must" version of `os.Truncate`
  - `func Unsetenv(key string)`: "must" version of `os.Unsetenv`
  - `func UserCacheDir() string`: "must" version of `os.UserCacheDir`
  - `func UserConfigDir() string`: "must" version of `os.UserConfigDir`
  - `func UserHomeDir() string`: "must" version of `os.UserHomeDir`
  - `func WriteFile(name string, data []byte, perm os.FileMode)`: "must" version of `os.WriteFile`
  - `type File`: "must" version of `os.File`
  - `type Process`: "must" version of `os.Process`
- fmtmust: "must" version of standard fmt package
  - `func Fprint(w Writer, a ...any) int`: "must" version of `fmt.Fprint`
  - `func Fprintf(w Writer, format string, a ...any) int`: "must" version of `fmt.Fprintf`
  - `func Fprintln(w Writer, a ...any) int`: "must" version of `fmt.Fprintln`
  - `func Fscan(r Reader, a ...any) int`: "must" version of `fmt.Fscan`
  - `func Fscanf(r Reader, format string, a ...any) int`: "must" version of `fmt.Fscanf`
  - `func Fscanln(r Reader, a ...any) int`: "must" version of `fmt.Fscanln`
  - `func Scan(a ...any) int`: "must" version of `fmt.Scan`
  - `func Scanf(format string, a ...any) int`: "must" version of `fmt.Scanf`
  - `func Scanln(a ...any) int`: "must" version of `fmt.Scanln`
  - `func Sscan(str string, a ...any) int`: "must" version of `fmt.Sscan`
  - `func Sscanf(str string, format string, a ...any) int`: "must" version of `fmt.Sscanf`
  - `func Sscanln(str string, a ...any) int`: "must" version of `fmt.Sscanln`
- timemust: "must" version of standard time package
  - `func LoadLocation(name string) *time.Location`: "must" version of `time.LoadLocation`
  - `func LoadLocationFromTZData(name string, data []byte) *time.Location`: "must" version of `time.LoadLocationFromTZData`
  - `func Parse(layout, value string) time.Time`: "must" version of `time.Parse`
  - `func ParseInLocation(layout, value string, loc *time.Location) time.Time`: "must" version of `time.ParseInLocation`
- pathmust/filepathmust: "must" version of standard path/filepath package
  - `func Abs(path string) string`: "must" version of `filepath.Abs`
  - `func EvalSymlinks(path string) string`: "must" version of `filepath.EvalSymlinks`
  - `func Glob(pattern string) []string`: "must" version of `filepath.Glob`
  - `func Localize(path string) string`: "must" version of `filepath.Localize`
  - `func Match(pattern, name string) bool`: "must" version of `filepath.Match`
  - `func Rel(basepath, targpath string) string`: "must" version of `filepath.Rel`
  - `func Walk(root string, fn filepath.WalkFunc)`: "must" version of `filepath.Walk`
  - `func WalkDir(root string, fn fs.WalkDirFunc)`: "must" version of `filepath.WalkDir`
- encodingmust/jsonmust: "must" version of standard encoding/json package
  - `func Compact(dst *bytes.Buffer, src []byte)`: "must" version of `json.Compact`
  - `func Indent(dst *bytes.Buffer, src []byte, prefix, indent string)`: "must" version of `json.Indent`
  - `func Marshal(v any) []byte`: "must" version of `json.Marshal`
  - `func MarshalIndent(v any, prefix, indent string) []byte`: "must" version of `json.MarshalIndent`
  - `func Unmarshal(data []byte, v any)`: "must" version of `json.Unmarshal`
  - `type Decoder`: "must" version of `json.Decoder`
  - `type Encoder`: "must" version of `json.Encoder`
- encodingmust/csvmust: "must" version of standard encoding/csv package
  - `type Reader`: "must" version of `csv.Reader`
  - `type Writer`: "must" version of `csv.Writer`
- encodingmust/base64must: "must" version of standard encoding/base64 package
  - `func NewDecoder(enc *base64.Encoding, r Reader) Reader`: "must" version of `base64.NewDecoder`
  - `func NewEncoder(enc *base64.Encoding, w Writer) WriteCloser`: "must" version of `base64.NewEncoder`
  - `type Encoding`: "must" version of `base64.Encoding`

## Motivation

When writing shell scripts, `set -e` is a common practice that makes the script exit immediately if any command fails.
This "fail-fast" behavior is useful for scripting because errors typically indicate something went wrong and continuing execution would be meaningless or dangerous.

```bash
set -e  # Exit immediately if any command fails

data=$(cat config.json)
```

However, Go's explicit error handling may create significant friction when writing scripts:

```go
// Traditional Go - verbose for scripts
data, err := os.ReadFile("config.json")
if err != nil {
    log.Panic(err)
}
```

**go-mustd brings the simplicity of `set -e` to Go scripting:**

```go
// With go-mustd - clean and concise like shell scripts
data := osmust.ReadFile("config.json")
```
