# winact-telegraf-plugin
This [Telegraf](https://www.influxdata.com/time-series-platform/telegraf/) plugin gathers activity information (uptime, active program, screen locked) from a Windows machine.

- [winact-telegraf-plugin](#winact-telegraf-plugin)
  * [Installation](#installation)
    + [Download `.exe`](#download-exe)
    + [Build From Source](#build-from-source)
  * [Usage](#usage)
  * [Configuration](#configuration)
  * [Metrics](#metrics)
    + [Types](#types)
    + [Example Output](#example-output)
    + [Example Influx Query](#example-influx-query)
  * [Grafana Dashboard](#grafana-dashboard)
  * [Credits](#credits)
  * [License](#license)

---

<br />

## Installation
### Download `.exe`
Simply download `winact-telegraf-plugin.exe` from the [latest release](https://github.com/zachstence/winact-telegraf-plugin/releases).

### Build From Source
*These instructions assume you have Go installed and configured on your machine*

Clone the repository
```sh
git clone https://github.com/zachstence/winact-telegraf-plugin
cd winact-telegraf-plugin
```

Build the module into an executable
```sh
go build -o winact-telegraf-plugin.exe cmd/main.go
```

## Usage
Reference the executable and config in your `telegraf.conf` using the `execd` input
```toml
[[inputs.execd]]
  command = ["/path/to/winact-telegraf-plugin.exe"] # no config
```

More documentation on using Telegraf external plugins can be found [here](https://github.com/influxdata/telegraf/blob/master/docs/EXTERNAL_PLUGINS.md).

## Configuration
At this time, no configuration options are available. If you have a need for one, feel free to open up an issue or PR!

```toml @sample.conf
[[inputs.winact]]
  # no config
```

## Metrics
TODO

### Types
- winact
  - TODO

### Example Output
TODO

### Example Influx Query
TODO

## Grafana Dashboard
Coming soon!

## Credits
TODO

## License
This project is subject to the the MIT License. See [LICENSE](./LICENSE) information for details.
