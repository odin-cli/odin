# odin

Create projects by remote templates, local templates,  or public repos

## Install

```shell
go install github.com/odin-cli/odin/cmd/cli/odin@latest
```

## Configure

To start, you need to configure Odin in your terminal. Use the following command:

```bash
odin config
```

The result:

```yaml
templates:
    default:
        repo: https://github.com/lbernardo/fx-webserver-example
        config:
            params:
                - name: package
            actions:
                - action: replace
                  find: github.com/lbernardo/fx-webserver
                  replace: '{{ .package }}'
```

## Create new project

Use following command:

```shell
odin create <NAME>
```

You can specifically template:

```shell
odin create <NAME> --template <TEMPLATE NAME, REPO ONLINE, TEMPLATE LOCAL> 
```

## Configuration odin cli

```yaml
# Example $HOME/.odin.config.yaml
templates:
    default: # Use: odin create <NAME> ('default' is default template) 
        repo: https://github.com/lbernardo/fx-webserver-example
        config:
            params:
                - name: package
            actions:
                - action: replace
                  find: github.com/lbernardo/fx-webserver
                  replace: '{{ .package }}'
    golang: # Use: odin create <NAME> --template golang
        local: /home/myuser/workspace/templates/golang
    front-end: # Use: odin create <NAME> --template front-end
        repo: https://github.com/myuser/myrepo-front-odin
```

| Param                    | Type                  | Required | Description                                                                                                         |
|--------------------------|-----------------------|----------|---------------------------------------------------------------------------------------------------------------------|
 | `template`               | `map[string]Template` | `no`     | Define templates                                                                                                    |
 | `template.<NAME>`        | `Template`            | `no`     | Template configurations                                                                                             |
 | `template.<NAME>.repo`   | `string`              | `no`     | When used, odin get template running: ```git clone <repo>```                                                        |
 | `template.<NAME>.local`  | `string`              | `no`     | When used, odin get template by path local                                                                          |
  | `template.<NAME>.config` | `ProjectConfig`       | `no`     | When using remote Repo's you can create project configurations if the repo doesn't have the .odin.yaml file in root |
  

## Project Config

You can use the project config to define actions to be taken when you run `create` with the specified template

There are two ways to create the Project Config:

1 - Creating file .odin.yaml in your project (repo, local)
```yaml
# .odin.yaml in your repo
apiVersion: v1
params:
  - name: package
actions:
  - action: replace
    find: github.com/lbernardo/fx-webserver
    replace: "{{ .package }}"
  - action: replace
    find: "${name}"
    replace: "{{ .name }}"
```


2 - Creating the configuration in config template file (`$HOME/.odin.config.yaml`)

```yaml
# $HOME/.odin.config.yaml
templates:
 fx-webserver:
  repo: https://github.com/lbernardo/fx-webserver-example
  config:
   apiVersion: "v1"
   params:
    - name: package
    - name: chart
    - name: repo
   actions:
    - action: replace
      find: github.com/lbernardo/fx-webserver
      replace: '{{ .package }}'
```

### Params

Use the params for replaces values in your project

Ex:

```shell
odin create teste --template fx-webserver
  INFO    Remote template
  INFO    Clonning project
  Cloning into 'teste'...
  remote: Enumerating objects: 29, done.
  remote: Counting objects: 100% (29/29), done.
  remote: Compressing objects: 100% (18/18), done.
  remote: Total 29 (delta 0), reused 29 (delta 0), pack-reused 0
  Receiving objects: 100% (29/29), 22.79 KiB | 2.53 MiB/s, done.
  
  Set value to 'package': ....
  Set value to 'chart': ....
  Set value to 'repo': .....
```

After that, you can use params to actions

### Actions

#### Replace

Use this action to execute value replaces in your new project

Ex:

```yaml
apiVersion: "v1"
params:
 - name: package
actions:
 - action: replace
   find: github.com/lbernardo/fx-webserver
   replace: '{{ .package }}'
```

In this action, odin find `github.com/lbernardo/fx-webserver` and replace by `{{ .package }}`

We use Go Template in value `replace:`, so, you can use for example: 

```yaml
apiVersion: "v1"
params:
 - name: package
actions:
 - action: replace
   find: github.com/lbernardo/fx-webserver
   replace: 'github.com/lbernardo/{{ .package }}'
```