package ProcessManager

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type ProcessStartConfig  struct {
	Directory string
	EntryFile string
	Enviroment string

}
type PermitedUser struct{
	User string
	ID   string
}
type FirstTimeStartConfig struct {
	Name string `json:"Name"`
	Owner string `json:"Owner"`
	ActiveDir string `json:"ActiveDir" `
	Token string  	`json:"Token"`
	PermitedUsers []PermitedUser `json:"PermittedUsers"`


}
func FirstTime_start(config FirstTimeStartConfig){
	first_config, err := os.Create(config.Name + ".json")
	if err!= nil {
        fmt.Println("Error creating file", err)
    }
	
	UserSettingsInit := FirstTimeStartConfig{
		Name: config.Name,
        Owner: config.Owner,
        ActiveDir: config.ActiveDir,
        Token: config.Token,
		
		
        
	}
	configSuccess, err := json.Marshal(UserSettingsInit)
	first_config.Write(configSuccess)
	first_config.Close()
	if err != nil {
		fmt.Println("Error encoding json", err)
	}
	

}

// func Start(start_set ProcessStartConfig) string{
	
// 	cmd := exec.Command("cd", start_set.Directory, "&&", start_set.Command)
//     cmd.Dir = start_set.Directory
//     out, err := cmd.CombinedOutput()
// 	if start_set.Command == ""{
// 		return "No command provided"
// 	}
// 	if start_set.Directory == ""{
// 		return "No directory provided"
// 	}

//     if err != nil {
//         fmt.Printf("cmd.Run() failed with error: %v\n", err)
        
//     }
	
// 	return string(out)
    


	
// }

func Stop(processName string) string{
	// "pm2", "stop", "PID"
    os.Chdir("procex/pm2-control/NodeEnviroment")
	switch processName{
	case "":
		return "No process name provided"
	}
    cmd := exec.Command("pm2", "stop", processName)
    out, err := cmd.CombinedOutput()
    if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
        
    }

    return string(out)

}

func Fetch_workspaces() string{
	
	Projects := exec.Command("cmd", "/C","cd .. && cd ActiveEnviroment && dir /ad /b")
	out, err := Projects.CombinedOutput()
	if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
        
    }
	
	return string(out)

	


}
func Delete(){

}



func List() string{
	
	// "pm2", "list"
	os.Chdir("procex/pm2-control/NodeEnviroment")
	cmd := exec.Command("pm2","list")
	out, err := cmd.CombinedOutput()
	if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
        
	}
	return string(out)

}

func Auth(authorize PermitedUser){

}

func  Init_project(){

}
func Pull(){


}
type Process struct{
	Name string
}
func Restart(process Process) string{
	// "pm2", "restart", "PID"
    os.Chdir("procex/pm2-control/NodeEnviroment")
    cmd := exec.Command("pm2", "restart", process.Name)
    out, err := cmd.CombinedOutput()
    if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
        
    }

    return string(out)

}