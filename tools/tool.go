package tools

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// somePath/APK_NAME/TOOL_NAME/DATA
var (
	ProcessedData string = "ProcessedData"

)

type file string


func NewTool(t string, f string) *Tool  {

	return &Tool {
		cmd: newCommand(t,f,filepath.Join(ProcessedData,f,t,f+"_src")),
		tool : t,
		file : f,
		dir: filepath.Join(ProcessedData,f,t,f+"_src"),
	}
}

type Tool struct {
	cmd *exec.Cmd
	tool string
	file string
	dir string
}





func  newCommand(t string,f string, d string) *exec.Cmd{

	cmdStruct := &exec.Cmd{}

	switch t {
	case "enjarify" :
		fmt.Println("Case Enjarify")
		// enjarify apk/filename
		// args := "d " + filepath.Join("..", "apk", f) + " -o " + SRC_DIR
		// args = filepath.Join("..","apk",t.file)
		args := []string{ "-f",  filepath.Join("..","..","..","..","apk",f)} //, " -o ",  "../" + d + f + ".jar"
		cmdStruct = exec.Command("enjarify",args...)

		// d := filepath.Join(ProcessedData,f,t)
		if err := os.MkdirAll(d,0666); err != nil {
			fmt.Println(err)
		}
		cmdStruct.Dir = d
	
	default :
		return nil
	}

	return cmdStruct
}


// @TODO enjarify should be able to input both apk and dex files
func  enjarify(f string, w http.ResponseWriter) {


}

// func (t *Tool) setDir()  string{
// 	return filepath.Join(ProcessedData, t.tool, t.file)
// }

func (t *Tool) Execute() {
	
	t.cmd.Stdout = os.Stdout
	t.cmd.Stderr = os.Stderr
	// t.cmd.Dir = t.setDir()

	fmt.Println(t.cmd.Args)
	fmt.Println(os.Getwd())
	err := t.cmd.Start()
	if err != nil {
		fmt.Println("Cannot Start Enjarify:",err)
	}

	err = t.cmd.Wait()
	if err != nil {
		fmt.Println("enjarify completion error", err)
	}
	fmt.Println("Reached End of Command")

}
