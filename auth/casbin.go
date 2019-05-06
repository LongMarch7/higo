package auth

import (
    "github.com/casbin/casbin"
    _ "github.com/go-sql-driver/mysql"
)
type Casbin struct {
    opts     authOpt
    enforcer *casbin.Enforcer
    adapter  *Adapter
}

func defaultConfig() authOpt{
    return authOpt{
        driverName:         "mysql",
        dataSourceName:     "root:123456@tcp(127.0.0.1:13306)/",
        baseName:           "test",
        dbSpecified:        false,
        ruleText:
`[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`,
    }
}

func NewCasbin(opts ...AOption) *Casbin {
    opt := defaultConfig()
    for _, o := range opts {
        o(&opt)
    }
    a := NewAdapter(opt.driverName,opt.dataSourceName,opt.baseName,opt.dbSpecified)
    m := casbin.NewModel(opt.ruleText)
    cas := &Casbin{
        opts: opt,
        adapter: a,
        enforcer: casbin.NewEnforcer(m, a),
    }
    cas.enforcer.LoadPolicy()
    return cas
}

func (c* Casbin)Enforcer() *casbin.Enforcer{
    return c.enforcer
}
