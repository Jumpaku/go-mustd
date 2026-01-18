package execmust

import (
	"context"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/Jumpaku/go-mustd"
	"github.com/Jumpaku/go-mustd/iomust"
	"github.com/Jumpaku/go-mustd/osmust"
)

func LookPath(file string) string {
	return mustd.Must1(exec.LookPath(file))
}
func Command(name string, arg ...string) *Cmd {
	return &Cmd{cmd: exec.Command(name, arg...)}
}
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	return &Cmd{cmd: exec.CommandContext(ctx, name, arg...)}
}

type Cmd struct {
	cmd *exec.Cmd
}

func (c *Cmd) SetPath(path string) {
	c.cmd.Path = path
}

func (c *Cmd) Path() string {
	return c.cmd.Path
}

func (c *Cmd) SetArgs(args []string) {
	c.cmd.Args = args
}

func (c *Cmd) Args() []string {
	return c.cmd.Args
}

func (c *Cmd) SetEnv(env []string) {
	c.cmd.Env = env
}

func (c *Cmd) Env() []string {
	return c.cmd.Env
}

func (c *Cmd) SetDir(dir string) {
	c.cmd.Dir = dir
}

func (c *Cmd) Dir() string {
	return c.cmd.Dir
}

func (c *Cmd) SetStdin(r iomust.Reader) {
	c.cmd.Stdin = r.Reader()
}

func (c *Cmd) Stdin() iomust.Reader {
	return iomust.ReaderOf(c.cmd.Stdin)
}

func (c *Cmd) SetStdout(r iomust.Writer) {
	c.cmd.Stdout = r.Writer()
}

func (c *Cmd) Stdout() iomust.Writer {
	return iomust.WriterOf(c.cmd.Stdout)
}

func (c *Cmd) SetStderr(r iomust.Writer) {
	c.cmd.Stderr = r.Writer()
}

func (c *Cmd) Stderr() iomust.Writer {
	return iomust.WriterOf(c.cmd.Stderr)
}

func (c *Cmd) SetExtraFiles(extraFiles []*osmust.File) {
	vs := make([]*os.File, len(extraFiles))
	for i, v := range extraFiles {
		vs[i] = v.File()
	}
	c.cmd.ExtraFiles = vs
}

func (c *Cmd) ExtraFiles() []*osmust.File {
	vs := make([]*osmust.File, len(c.cmd.ExtraFiles))
	for i, v := range c.cmd.ExtraFiles {
		vs[i] = osmust.FileOf(v)
	}
	return vs
}

func (c *Cmd) SetSysProcAttr(sysProcAttr *syscall.SysProcAttr) {
	c.cmd.SysProcAttr = sysProcAttr
}

func (c *Cmd) SysProcAttr() *syscall.SysProcAttr {
	return c.cmd.SysProcAttr
}

func (c *Cmd) SetProcess(process *os.Process) {
	c.cmd.Process = process
}

func (c *Cmd) Process() *os.Process {
	return c.cmd.Process
}

func (c *Cmd) SetProcessState(processState *os.ProcessState) {
	c.cmd.ProcessState = processState
}

func (c *Cmd) ProcessState() *os.ProcessState {
	return c.cmd.ProcessState
}

func (c *Cmd) SetErr(err error) {
	c.cmd.Err = err
}

func (c *Cmd) Err() error {
	return c.cmd.Err
}

func (c *Cmd) SetCancel(cancel func() error) {
	c.cmd.Cancel = cancel
}

func (c *Cmd) Cancel() func() error {
	return c.cmd.Cancel
}

func (c *Cmd) SetWaitDelay(waitDelay time.Duration) {
	c.cmd.WaitDelay = waitDelay
}

func (c *Cmd) WaitDelay() time.Duration {
	return c.cmd.WaitDelay
}

func (c *Cmd) CombinedOutput() []byte {
	return mustd.Must1(c.cmd.CombinedOutput())
}
func (c *Cmd) Environ() []string {
	return c.cmd.Environ()
}
func (c *Cmd) Output() []byte {
	return mustd.Must1(c.cmd.Output())
}
func (c *Cmd) Run() {
	mustd.Must0(c.cmd.Run())
}
func (c *Cmd) Start() {
	mustd.Must0(c.cmd.Start())
}
func (c *Cmd) StderrPipe() iomust.ReadCloser {
	return iomust.ReadCloserOf(mustd.Must1(c.cmd.StderrPipe()))
}
func (c *Cmd) StdinPipe() iomust.WriteCloser {
	return iomust.WriteCloserOf(mustd.Must1(c.cmd.StdinPipe()))
}
func (c *Cmd) StdoutPipe() iomust.ReadCloser {
	return iomust.ReadCloserOf(mustd.Must1(c.cmd.StdoutPipe()))
}
func (c *Cmd) String() string {
	return c.cmd.String()
}
func (c *Cmd) Wait() {
	mustd.Must0(c.cmd.Wait())
}
