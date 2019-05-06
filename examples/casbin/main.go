package main

import (
    "fmt"
    "github.com/LongMarch7/higo/auth"
)

func main() {
    cas :=auth.NewCasbin()

    police := []string{"admin", "data1", "read"}
    cas.Enforcer().AddPolicy(police)
    police1 := []string{"user", "data2", "read"}
    cas.Enforcer().AddPolicy(police1)
    cas.Enforcer().AddRoleForUser("alice","admin")
    // Check the permission.
    if cas.Enforcer().Enforce("alice", "data1", "read") {
        fmt.Println("allow")
    }else{
        fmt.Println("deny")
    }
    if cas.Enforcer().Enforce("alice", "data2", "read") {
        fmt.Println("allow")
    }else{
        fmt.Println("deny")
    }

    cas.Enforcer().AddRoleForUser("petter","alice")
    cas.Enforcer().AddRoleForUser("petter","user")
    // Check the permission.
    if cas.Enforcer().Enforce("petter", "data1", "read") {
        fmt.Println("allow")
    }else{
        fmt.Println("deny")
    }
    if cas.Enforcer().Enforce("petter", "data2", "read") {
        fmt.Println("allow")
    }else{
        fmt.Println("deny")
    }

    // Modify the policy.
    // e.AddPolicy(...)
    // e.RemovePolicy(...)

    // Save the policy back to DB.
    //e.SavePolicy()

}