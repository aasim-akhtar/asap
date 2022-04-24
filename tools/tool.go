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


func NewTool(toolName string, fileName string) *Tool  {

	outputDir := filepath.Join(ProcessedData,fileName,toolName)
	return &Tool {
		cmd: newCommand(toolName,fileName,outputDir),
		tool : toolName,
		file : fileName,
		dir: outputDir,
	}
}

type Tool struct {
	cmd *exec.Cmd
	tool string
	file string
	dir string
}





func  newCommand(t string,f string, d string) *exec.Cmd{

	var cmdStruct *exec.Cmd
	apkPath :=  filepath.Join("..","..","..","apk",f)
	Mkdir(d)

	switch t {
	case "enjarify" :
		fmt.Println("Case Enjarify")
		// enjarify apk/filename
		// args := "d " + filepath.Join("..", "apk", f) + " -o " + SRC_DIR
		// args = filepath.Join("..","apk",t.file)
		args := []string{ "-f", apkPath} //, " -o ",  "../" + d + f + ".jar"
		cmdStruct = exec.Command("enjarify",args...)
	
	case "jadx" :
		fmt.Println("Case jadx")
		// jadx -d ../../../ProcessedData/out ../../../apk/filexyz.apk
		// or
		// jadx ../../../apk/file.apk
		args := []string{apkPath}

		cmdStruct = exec.Command("jadx",args...)

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
	t.cmd.Dir = t.dir
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

// Mkdir makes Output directory for tool to be executed.
// If output directory is not already present some tools may fail to execute
func Mkdir(dir string) {
	if err := os.MkdirAll(dir,0666); err != nil {
		fmt.Println(err)
	}
}