package osmust

import (
	"os"
	"time"

	"github.com/Jumpaku/go-mustd"
)

// Chdir changes the current working directory. Panics if an error occurs.
func Chdir(dir string) {
	mustd.Must0(os.Chdir(dir))
}

// Chmod changes the mode of the named file. Panics if an error occurs.
func Chmod(name string, mode os.FileMode) {
	mustd.Must0(os.Chmod(name, mode))
}

// Chown changes the numeric uid and gid of the named file. Panics if an error occurs.
func Chown(name string, uid, gid int) {
	mustd.Must0(os.Chown(name, uid, gid))
}

// Chtimes changes the access and modification times of the named file. Panics if an error occurs.
func Chtimes(name string, atime time.Time, mtime time.Time) {
	mustd.Must0(os.Chtimes(name, atime, mtime))
}

// Executable returns the path name for the executable that started the current process. Panics if an error occurs.
func Executable() string {
	return mustd.Must1(os.Executable())
}

// Getgroups returns a list of the numeric ids of groups that the caller belongs to. Panics if an error occurs.
func Getgroups() []int {
	return mustd.Must1(os.Getgroups())
}

// Getwd returns the current working directory. Panics if an error occurs.
func Getwd() (dir string) {
	return mustd.Must1(os.Getwd())
}

// Hostname returns the host name. Panics if an error occurs.
func Hostname() (name string) {
	return mustd.Must1(os.Hostname())
}

// Lchown changes the numeric uid and gid of the named file without following symbolic links. Panics if an error occurs.
func Lchown(name string, uid, gid int) {
	mustd.Must0(os.Lchown(name, uid, gid))
}

// Link creates newname as a hard link to the oldname file. Panics if an error occurs.
func Link(oldname, newname string) {
	mustd.Must0(os.Link(oldname, newname))
}

// Mkdir creates a new directory with the specified name and permission bits. Panics if an error occurs.
func Mkdir(name string, perm os.FileMode) {
	mustd.Must0(os.Mkdir(name, perm))
}

// MkdirAll creates a directory named path, along with any necessary parents. Panics if an error occurs.
func MkdirAll(path string, perm os.FileMode) {
	mustd.Must0(os.MkdirAll(path, perm))
}

// MkdirTemp creates a new temporary directory in the directory dir. Panics if an error occurs.
func MkdirTemp(dir, pattern string) string {
	return mustd.Must1(os.MkdirTemp(dir, pattern))
}

// Pipe returns a connected pair of Files. Panics if an error occurs.
func Pipe() (r *File, w *File) {
	rf, wf := mustd.Must2(os.Pipe())
	return FileOf(rf), FileOf(wf)
}

// ReadFile reads the named file and returns the contents. Panics if an error occurs.
func ReadFile(name string) []byte {
	return mustd.Must1(os.ReadFile(name))
}

// Readlink returns the destination of the named symbolic link. Panics if an error occurs.
func Readlink(name string) string {
	return mustd.Must1(os.Readlink(name))
}

// Remove removes the named file or empty directory. Panics if an error occurs.
func Remove(name string) {
	mustd.Must0(os.Remove(name))
}

// RemoveAll removes path and any children it contains. Panics if an error occurs.
func RemoveAll(path string) {
	mustd.Must0(os.RemoveAll(path))
}

// Rename renames (moves) oldpath to newpath. Panics if an error occurs.
func Rename(oldpath, newpath string) {
	mustd.Must0(os.Rename(oldpath, newpath))
}

// Setenv sets the value of the environment variable named by the key. Panics if an error occurs.
func Setenv(key, value string) {
	mustd.Must0(os.Setenv(key, value))
}

// Symlink creates newname as a symbolic link to oldname. Panics if an error occurs.
func Symlink(oldname, newname string) {
	mustd.Must0(os.Symlink(oldname, newname))
}

// Truncate changes the size of the named file. Panics if an error occurs.
func Truncate(name string, size int64) {
	mustd.Must0(os.Truncate(name, size))
}

// Unsetenv unsets the specified environment variable. Panics if an error occurs.
func Unsetenv(key string) {
	mustd.Must0(os.Unsetenv(key))
}

// UserCacheDir returns the default root directory to use for user-specific cached data. Panics if an error occurs.
func UserCacheDir() string {
	return mustd.Must1(os.UserCacheDir())
}

// UserConfigDir returns the default root directory to use for user-specific configuration data. Panics if an error occurs.
func UserConfigDir() string {
	return mustd.Must1(os.UserConfigDir())
}

// UserHomeDir returns the current user's home directory. Panics if an error occurs.
func UserHomeDir() string {
	return mustd.Must1(os.UserHomeDir())
}

// WriteFile writes data to the named file, creating it if necessary. Panics if an error occurs.
func WriteFile(name string, data []byte, perm os.FileMode) {
	mustd.Must0(os.WriteFile(name, data, perm))
}

// Create creates or truncates the named file. Panics if an error occurs.
func Create(name string) *File {
	return &File{file: mustd.Must1(os.Create(name))}
}

// CreateTemp creates a new temporary file in the directory dir. Panics if an error occurs.
func CreateTemp(dir, pattern string) *File {
	return &File{file: mustd.Must1(os.CreateTemp(dir, pattern))}
}

// Open opens the named file for reading. Panics if an error occurs.
func Open(name string) *File {
	return &File{file: mustd.Must1(os.Open(name))}
}

// OpenFile opens the named file with specified flag and perm. Panics if an error occurs.
func OpenFile(name string, flag int, perm os.FileMode) *File {
	return &File{file: mustd.Must1(os.OpenFile(name, flag, perm))}
}

// Lstat returns a FileInfo describing the named file without following symbolic links. Panics if an error occurs.
func Lstat(name string) os.FileInfo {
	return mustd.Must1(os.Lstat(name))
}

// Stat returns a FileInfo describing the named file. Panics if an error occurs.
func Stat(name string) os.FileInfo {
	return mustd.Must1(os.Stat(name))
}

// Process wraps os.Process and provides panicking error handling.
type Process struct {
	process *os.Process
}

// FindProcess looks for a running process by its pid. Panics if an error occurs.
func FindProcess(pid int) *Process {
	return &Process{process: mustd.Must1(os.FindProcess(pid))}
}

// StartProcess starts a new process with the program, arguments, and attributes. Panics if an error occurs.
func StartProcess(name string, argv []string, attr *os.ProcAttr) *Process {
	return &Process{process: mustd.Must1(os.StartProcess(name, argv, attr))}
}

// Kill causes the Process to exit immediately. Panics if an error occurs.
func (p *Process) Kill() {
	mustd.Must0(p.process.Kill())
}

// Release releases any resources associated with the Process. Panics if an error occurs.
func (p *Process) Release() {
	mustd.Must0(p.process.Release())
}

// Signal sends a signal to the Process. Panics if an error occurs.
func (p *Process) Signal(sig os.Signal) {
	mustd.Must0(p.process.Signal(sig))
}

// Wait waits for the Process to exit. Panics if an error occurs.
func (p *Process) Wait() *os.ProcessState {
	return mustd.Must1(p.process.Wait())
}
