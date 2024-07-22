package cli

import (
	"os"
	"strings"
)

type CLI struct {
	command            string
	resources          []string
	args               map[string][]string
	tasks              map[string]Task
	resourcesCallbacks []interface{}
	errors             []interface{}
	init               []interface{}
	shutdown           []interface{}
}

func NewCLI(args []string) (cli CLI) {
	cli = CLI{}

	if args == nil || len(args) == 0 {
		args = os.Args
	}

	cli.args = cli.Parse(args)

	return
}

func (cli *CLI) Init() {
	cli.init = make([]interface{}, 0)
}

func (cli *CLI) Shutdown() {
	cli.shutdown = make([]interface{}, 0)
}

func (cli *CLI) Error() {
	cli.errors = make([]interface{}, 0)
}

func (cli *CLI) Task(name string) (task Task) {
	task = Task{Name: name}
	cli.tasks[name] = task

	return
}

func (cli *CLI) Parse(args []string) map[string][]string {
	args = args[1:] // Remove script path

	if len(args) == 0 {
		panic("Missing command")
	}

	cli.command = args[0]

	args = args[1:]
	if len(args) == 0 {
		return nil
	}

	output := make(map[string][]string)

	for i, arg := range args {
		if strings.Index(arg, "--") == 0 {
			args[i] = arg[2:]
		}
	}

	for _, arg := range args {
		pair := strings.Split(arg, "=")
		if len(pair) == 2 {
			output[pair[0]] = append(output[pair[0]], pair[1])
		}
	}

	return output
}
