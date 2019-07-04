package shell

import (
    "os/exec"
    "runtime"
)

//command是执行的shell,注意如果出错，string(out)的报错内容更详细
func Exec_Shell(command string) (string, error) {
    sysType := runtime.GOOS
    var cmd *exec.Cmd

    if sysType == "linux"{
        cmd = exec.Command("/bin/sh", "-c", command)
    }else if sysType == "windows" {
        cmd = exec.Command("cmd", "/c", command)
    }
    out, err := cmd.CombinedOutput()
    if err !=nil{
        return "", err
    }
    return string(out), err
}
