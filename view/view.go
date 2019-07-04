package view

import (
    "fmt"
    "google.golang.org/grpc/grpclog"
    "html/template"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "sync"
)

// View is an interface for rendering templates.
type View interface {
    Render(out io.Writer, name string, data interface{}) error
}

// SimpleView implements View interface, but based on golang templates.
type ViewStruct struct {
    tmpl      *template.Template
    funcMap   template.FuncMap
}

var onceAction sync.Once
var templator View
func defaultConfig() viewOpt{
    return viewOpt{
        dir: "E:/go_project/higo/src/github.com/LongMarch7/higo/template",
        delimsLeft: "{{",
        delimsRight: "}}",
    }
}
func NewView(opts ...VOption) View{
    onceAction.Do(func() {
        opt := defaultConfig()
        for _, o := range opts {
            o(&opt)
        }
        template, err := viewInit(opt)
        if err != nil{
            panic(err)
        }
        templator = template
    })
    return templator
}
//NewSimpleView returns a SimpleView with templates loaded from viewDir
func viewInit(opt viewOpt) (View, error) {
    info, err := os.Stat(opt.dir)
    if err != nil {
        return nil, err
    }
    if !info.IsDir() {
        return nil, fmt.Errorf("utron: %s is not a directory", opt.dir)
    }
    s := &ViewStruct{
        tmpl:    template.New(filepath.Base(opt.dir)),
        funcMap: make(template.FuncMap),
    }
    s.funcMap["html2str"] = HTML2str
    s.funcMap["str2html"] = Str2html
    s.funcMap["urlfor"] = URLFor
    s.funcMap["convertm"] = ConvertM
    s.funcMap["convertt"] = ConvertT
    s.tmpl.Delims(opt.delimsLeft,opt.delimsRight)
    s.tmpl.Funcs(s.funcMap)
    //s.tmpl.Delims("{<",">}")
    return s.load(opt.dir)
}

// load loads templates from dir. The templates should be valid golang templates
//
// Only files with extension .html, .tpl, .tmpl will be loaded. references to these templates
// should be relative to the dir. That is, if  dir is foo, you don't have to refer to
// foo/bar.tpl, instead just use bar.tpl
func (s *ViewStruct) load(dir string) (View, error) {

    // supported is the list of file extensions that will be parsed as templates
    supported := map[string]bool{".tpl": true, ".html": true, ".tmpl": true}

    werr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }

        extension := filepath.Ext(path)
        if _, ok := supported[extension]; !ok {
            return nil
        }

        data, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }

        // We remove the directory name from the path
        // this means if we have directory foo, with file bar.tpl
        // full path for bar file foo/bar.tpl
        // we trim the foo part and remain with /bar.tpl
        //
        // NOTE we don't account for the opening slash, when dir ends with /.
        name := path[len(dir):]

        name = filepath.ToSlash(name)

        name = strings.TrimPrefix(name, "/") // case  we missed the opening slash

        name = strings.TrimSuffix(name, extension) // remove extension
        grpclog.Info(name)
        t := s.tmpl.New(name)
        if _, err = t.Parse(string(data)); err != nil {
            return err
        }
        return nil
    })

    if werr != nil {
        return nil, werr
    }

    return s, nil
}

// Render executes template named name, passing data as context, the output is written to out.
func (s *ViewStruct) Render(out io.Writer, name string, data interface{}) error {
    return s.tmpl.ExecuteTemplate(out, name, data)
}
