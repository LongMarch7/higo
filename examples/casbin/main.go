package main

import (
    "fmt"
    "github.com/LongMarch7/higo/auth"
    "github.com/LongMarch7/higo/db"
)

func main() {
    db.NewDb(db.DefaultNAME, db.Dialect("mysql"),
        db.Args("root:123456@tcp(localhost:13306)/higo?charset=utf8"),
        db.MaxOpenConns(100),
        db.MaxIdleConns(100))
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
    cas.Enforcer().AddRoleForUser("john","petter")
    cas.Enforcer().AddRoleForUser("kaite","petter")
    fmt.Println(cas.Enforcer().GetUsersForRole("petter"))
    // Check the permission.
    if cas.Enforcer().Enforce("kaite", "data1", "read") {
        fmt.Println("allow")
    }else{
        fmt.Println("deny")
    }
    if cas.Enforcer().Enforce("petter", "data2", "read") {
        fmt.Println("allow")
    }else{
        fmt.Println("deny")
    }
    //cas.Enforcer().RemovePolicy(police1)
    //cas.Enforcer().RemoveGroupingPolicy("petter","alice")
    // Modify the policy.
    // e.AddPolicy(...)
    // e.RemovePolicy(...)

    // Save the policy back to DB.
    //cas.Enforcer().SavePolicy()

}