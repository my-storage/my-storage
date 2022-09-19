# my-storage

#### Folder structure

```sh
infrastructure # All infrastructure files for non-develop environment, Terraform, Kubernetes, etc.
ms-*
  src
   ├── main         # Configures, startup, instantiatio and observability
   ├── presentation # All files to comunication aync/async.
   │      ├──  rest
   │      ├──  grpc
   │      └──  pub sub
   ├── shared       # All shared stuphs
   │      ├──  protocols
   │      ├──  data models
   │      ├──  base adapters
   │      ├──  agregators
   │      └──  utils
   ├── modules      # Business rules layers
   │      └── example-module
   │                ├──  usecases
   │                ├──  repositories
   │                └──  shared # Same from above but only the necessary
```

This structure is only a reference to the folder implementation.
