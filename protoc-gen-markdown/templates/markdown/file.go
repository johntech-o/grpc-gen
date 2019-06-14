package markdown

const fileTpl = `
# {{ .InputPath }} 接口文档

*Document generated by protoc-gen-markdown. DO NOT EDIT.*

> 接口列表

{{ range $v := (toc .) }}
{{ if $v.Interface }}
	* [{{ $v.Name.UpperCamelCase }}{{ if $v.Gateway }} ({{ $v.Gateway }}){{ end }}](#{{ anchor $v.Name }}) - {{ $v.Comment }}
{{ else }}
* [{{ $v.Name.UpperCamelCase }}](#{{ anchor $v.Name }}) - {{ $v.Comment }}
{{ end }}
{{ end }}

{{ range .Services }}
{{ template "service" . }}
{{ end }}

## *Embed Messages*

{{ range $type, $m := (embedMessages .) }}
{{ if (eq $type "enum") }}

{{ range $name, $t := $m }}
<h3 id="{{ anchor $t.Name }}">{{ $t.Name.UpperCamelCase }}</h3>

> 枚举类型

|取值|对应数值|说明|
|---|---|---|
{{ range $v := $t.Values }}|{{ $v.Name }}|{{ $v.Value }}|{{ $v.Comment }}|
{{ end }}
{{ end }}

{{ else }}

{{ range $name, $t := $m }}
<h3 id="{{ anchor $t.Name }}">{{ $t.Name.UpperCamelCase }}</h3>

> 自定义类型

|字段|类型|说明|默认值|是否必传|
|---|---|---|---|---|
{{ range $v := $t.Fields }}|{{ $v.Name }}|{{ $v.Type }}|{{ $v.Comment }}|-|{{ $v.Required }}|
{{ end }}
{{ end }}

{{ end }}
{{ end }}
`