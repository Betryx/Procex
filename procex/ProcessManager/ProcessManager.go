package ProcessManager

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"github.com/google/uuid"
	
)

type ProcessStartConfig  struct {
	WorkSpace string
	ProjectName string
	RootDirectory string
	EntryFile string
	Enviroment string

}
type PermitedUser struct{
	UserName string
	ID   string
}
type WorkSpace struct {
	Name string `json:"Name"`
	Owner string `json:"Owner"`
	Token string  	`json:"Token"`
	PermitedUsers []PermitedUser `json:"PermittedUsers"`
	GitHubUserName string  `json:"GitHubUserName"`
	Uuid string  `json:"Uuid"`


}
func Start(start_set ProcessStartConfig) string{
	os.Chdir("../pm2-control/NodeEnviroment")
	cmd := exec.Command("pm2", "start", "../ActiveEnviroment/"+start_set.WorkSpace+"/"+start_set.ProjectName+"/"+start_set.RootDirectory+"/"+start_set.EntryFile, "--name "+start_set.ProjectName, "--interpreter", start_set.Enviroment)
	out, err := cmd.CombinedOutput()
	if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
        return string(out)
    }
	return string(out)
	
}





















func WorkSpaceCreate(config WorkSpace) string{
	
	os.Chdir("../ActiveEnviroment")
	
	os.MkdirAll(config.Name, 0777)
	first_config, err := os.Create(config.Name+".json")
	if err!= nil {
        fmt.Println("Error creating file", err)
		return string("Error creating Workspace " + config.Name + "with owner" + config.Owner)
    }
	
	UserSettingsInit := WorkSpace{
		Name: config.Name,
        Owner: config.Owner,
        Token: config.Token,
		PermitedUsers: []PermitedUser{},
		GitHubUserName: config.GitHubUserName,
		Uuid: uuid.New().String(),
		
		
        
	}
	configSuccess, err := json.Marshal(UserSettingsInit)
	first_config.Write(configSuccess)
	first_config.Close()
	if err != nil {
		fmt.Println("Error encoding json", err)
	}
	return string("WorkSpace "+config.Name +" created. Owner is " + config.Owner)
	

}


type ProcessStop struct {
	ProcessName string
	WorkSpace string
	Uuid string


}
func Stop(stop ProcessStop) string{
	os.Chdir("../ActiveEnviroment")
	var data WorkSpace
	file, err := os.ReadFile(stop.WorkSpace+".json")
	if err!= nil {
        fmt.Println("Error reading file", err)
        return "Error while stopping process"
    }
	json.Unmarshal(file, &data)
	if stop.Uuid != data.Uuid {
		return "Provided auth code is incorrect"
	}
	// "pm2", "stop", "PID"
	os.Chdir("../procex")
    os.Chdir("../pm2-control/NodeEnviroment")
	if stop.ProcessName == ""{
		return "No process name provided"
	}

    cmd := exec.Command("pm2", "stop", stop.ProcessName)
    out, err := cmd.CombinedOutput()
    if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
		return string(out)
        
    }

    return string("Process "+stop.ProcessName+" stopped successfully\n Output: "+string(out))

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
type ProjectInit struct {
	Name string
	Owner string
	Workspace string
	ProjectRepo string


}
func  Init_project(project * ProjectInit) string{
	os.Chdir("../ActiveEnviroment")
	file, err := os.ReadFile(project.Workspace+".json")
	if err!= nil {
        fmt.Println("Error reading file", err)
        return "Error while initializing project"
    }
	
	
	var data WorkSpace
	json.Unmarshal(file, &data)
	os.Chdir("../ActiveEnviroment/"+project.Workspace)
	os.MkdirAll(project.Name, 0777)
	os.Chdir(project.Name)
	template := "https://"+data.GitHubUserName+":"+data.Token+"/@github.com/"+data.GitHubUserName+"/"+project.Name+".git"
	fmt.Println(template)
	return template
	
}
func Pull(){


}
type Process struct{
	Name string
}
func Restart(process Process) string{
	// "pm2", "restart", "PID"
	if process.Name == ""{
		return "A process with the name provided either doesnt exist or the name is wrong"
	}
    os.Chdir("procex/pm2-control/NodeEnviroment")
    cmd := exec.Command("pm2", "restart", process.Name)
    out, err := cmd.CombinedOutput()
    if err!= nil {
        fmt.Printf("cmd.Run() failed with error: %v\n", err)
        
    }

    return string(out)

}

