# yml-expander
Create simple and easy yml templates without massive dependency instalation


## Instalation

Install a release from: [github.com/wpjunior/yml-expander/releases](github.com/wpjunior/yml-expander/releases)

## Example of usage

Create a context file named `data.yml`

```
environments:
   - dev
   - staging
```

Create a template file named config.yml.tpl

```
environments:
{{ range $_, $environment := .environments }}
- name: example.{{ $environment }}
  testing: true
{{ end }}
```

Generate `config.yml` running

```
yml-exapander -data.path=data.yml -template.path=config.yml.tpl -output.path=config.yml
```

The output will be:

```
environments:
- name: example.dev
  testing: true
- name: example.staging
  testing: true
```
