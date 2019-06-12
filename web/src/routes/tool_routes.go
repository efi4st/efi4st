package routes

import (
	"fmt"
	"github.com/kataras/iris"
	"os/exec"
)

func ToolRun(ctx iris.Context) {

	var (
		toolname = ctx.Params().Get("toolname")
	)

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println(toolname)

	cmd := exec.Command("job.sh")
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	fmt.Println(string(out))

	ctx.Redirect("/")
}