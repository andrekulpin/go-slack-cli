package main

import "fmt"
import "os/exec"
import "github.com/urfave/cli"


func check(err error){
	if(err != nil){
		panic(err)
	}
}

func main(){

	app := cli.NewApp();
	app.Name = "Slack";
	app.Version = "0.0.1";

}



func Login(){
	args := []string{}
	runCommand("curl", )
}

func runCommand(cmd string, args []string) ( []byte, error ) {
	out, res := exec.Command(cmd, args...).Output();
	return out, res;
}

/*type Response struct {
	Name string
	City string
	XForwardedFor string `json:x-forwarded-for`
}*/
