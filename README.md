# my-storage

#### Scripting

We using Taskfile in project so you need to install go module to use:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

References in https://taskfile.dev/installation/#go-modules.

#### Folder structure

```sh
infrastructure # All infrastructure files for non-develop environment, Terraform, Kubernetes, etc.
ms-*
  src
   ├── main.go      # Application start point
   ├── app          # Configures, startup and application layer (usecases, helpers and something like that)
   ├── modules      # Business rules layers
   │      └── example-module
   │                ├──  entities
   │                ├──  repositories
   │                └──  shared # Same from above but only the necessary
   ├── presentation # All communication files sync/async. Normally this was in shared/infra but because of we have a large file numbers we try a different approach here.
   │      ├──  rest
   │      ├──  grpc
   │      └──  pub sub
   ├── shared       # All shared stuphs
   │      ├──  aggregators
   │      ├──  infra
   │      ├──  data models
   │      ├──  protocols
   │      └──  utils
```

This structure is only a reference to the folder implementation.
