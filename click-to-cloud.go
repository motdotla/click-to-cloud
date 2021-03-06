package main

import (
  cli "github.com/codegangsta/cli"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  simplejson "github.com/bitly/go-simplejson"
  "strings"
  "strconv"
  "time"
)

func main() {
  app       := cli.NewApp()
  app.Name  = "click-to-cloud"
  app.Usage = "Deploy applications to the cloud with just a 'click'."

  app.Flags = []cli.Flag {
    cli.StringFlag{"repo", "", "Git url repository of click-to-cloud enabled app"},
  }

  app.Action = func(c *cli.Context) {
    nano_count  := strconv.FormatInt(time.Now().UnixNano(), 10)
    deploy_dir  := "deploy-" + nano_count
    fmt.Println(deploy_dir)

    os.MkdirAll(deploy_dir, 0777)
    os.Chdir(deploy_dir)

    repo_url        := c.String("repo")
    repo_url_split  := strings.Split(repo_url, "/")
    repo_url_split_length := len(repo_url_split)
    repo_name       := strings.Split(repo_url_split[repo_url_split_length-1], ".git")[0]

    // Step 1: Clone repo 
    clone_cmd         := exec.Command("git", "clone", repo_url)
    clone_cmd.Stdout  = os.Stdout
    clone_cmd.Stderr  = os.Stderr
    err               := clone_cmd.Run()
    if err != nil {
      log.Fatal(err)
    }

    // Step 2: chdir
    os.Chdir(repo_name)

    // Step 3: Read and parse click-to-cloud.json
    data, err := ioutil.ReadFile("./click-to-cloud.json")
    json, err := simplejson.NewJson(data)
    if err != nil {
      log.Fatal(err)
    }

    // Step 4: run commands
    commands, _ := json.Get("heroku").Array()

    for _, command := range commands {
      command_to_string, _ := command.(string)
      fmt.Println(command_to_string)
      fmt.Println("Running command: "+command_to_string)

      command_slice           := strings.Split(command_to_string, " ")
      command_slice_end_range := len(command_slice)

      exec_command        := exec.Command(command_slice[0], command_slice[1:command_slice_end_range] ...)
      exec_command.Stdout = os.Stdout
      exec_command.Stderr = os.Stderr
      err                 := exec_command.Run()
      if err != nil {
        log.Fatal(err)
      }
    }

    // Step 5: cleanup
    os.Chdir("..")
    os.Chdir("..")
    os.RemoveAll(deploy_dir)
  }

  app.Run(os.Args)
}
