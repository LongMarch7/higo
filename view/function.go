package view

import (
    "github.com/alioygur/godash"
    "html/template"
    "regexp"
    "strings"
    "time"
)

func Str2html(raw string) template.HTML {
    return template.HTML(raw)
}

func HTML2str(html string) string {
    re := regexp.MustCompile(`\<[\S\s]+?\>`)
    html = re.ReplaceAllStringFunc(html, strings.ToLower)

    //remove STYLE
    re = regexp.MustCompile(`\<style[\S\s]+?\</style\>`)
    html = re.ReplaceAllString(html, "")

    //remove SCRIPT
    re = regexp.MustCompile(`\<script[\S\s]+?\</script\>`)
    html = re.ReplaceAllString(html, "")

    re = regexp.MustCompile(`\<[\S\s]+?\>`)
    html = re.ReplaceAllString(html, "\n")

    re = regexp.MustCompile(`\s{2,}`)
    html = re.ReplaceAllString(html, "\n")

    return strings.TrimSpace(html)
}

func URLFor(url string, f string) string {
    return "/"+ url + "/" + strings.ToLower(godash.ToSnakeCase(f))
}

func ConvertT(in int) (out string) {
    tm := time.Unix(int64(in), 0)
    return tm.Format("2006-01-02 15:04:05")
}

func ConvertM(in int) (out int64) {
    return int64(in) / 1000000
}
