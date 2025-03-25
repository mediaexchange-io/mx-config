# appconfig

After a search for a *simple* application configuration library turned up many
complex libraries, we decided to publish our own.

## Features

* Data structure is defined by the application.
* Strongly typed configuration data (no key-value maps)
* File formats: JSON and YAML
* Overridden values read from environment variables.

## Usage

This small library solves a common use-case found in microservice applications
where configuration values come from a combination of files and environment
variables, e.g. Kubernetes ConfigMaps and Secrets.

As an example, let's say we need to configure the port number the web service
listens on and the connection information for a database. The YAML file could
look like this:

```YAML
---
server:
  port: 8080

database:
  driver:   postgres
  hostname: localhost
  port:     5432
  username: postgres
  password: dummy
  name:     my_database
```

Now in the Go application a data structure is defined to match:

```go
type AppConfig struct {
    Server struct {
        Port int16
    }
    Database struct {
        Driver   string
        Hostname string
        Port     int16
        Username string
        Password string `env:"DATABASE_PASSWORD"`
        Name     string
    }
}
```

Notice the database password has an environment variable associated with it.
With this, the actual password may be provided from, for example, a Kubernetes
secret that is passed through the environment rather than a mapped file.

To ingest the configuration file and any environment variable overrides into
the application is just two lines of code:

```go
var c = new(AppConfig)
var err = config.FromFile("/path/of/config.yaml", c)
```

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Media Exchange](http://github.com/mediaexchange/)

## License

   Copyright 2025 MediaExchange.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
