<h1 align="center">
  <a href="https://github.com/YojimboSecurity/skyplatform">
    <!-- Please provide path to your logo here -->
    <img src="docs/images/logo.png" alt="Logo" width="108" height="108">
  </a>
</h1>

<div align="center">
  Skyplatform
  <br />
  <br />
  <a href="https://github.com/YojimboSecurity/skyplatform/issues/new?assignees=&labels=bug&template=01_BUG_REPORT.md&title=bug%3A+">Report a Bug</a>
  ·
  <a href="https://github.com/YojimboSecurity/skyplatform/issues/new?assignees=&labels=enhancement&template=02_FEATURE_REQUEST.md&title=feat%3A+">Request a Feature</a>
  .
  <a href="https://github.com/YojimboSecurity/skyplatform/issues/new?assignees=&labels=question&template=04_SUPPORT_QUESTION.md&title=support%3A+">Ask a Question</a>
</div>

<div align="center">
<br />

[![Project license](https://img.shields.io/github/license/YojimboSecurity/skyplatform.svg?style=flat-square)](LICENSE)

[![Pull Requests welcome](https://img.shields.io/badge/PRs-welcome-ff69b4.svg?style=flat-square)](https://github.com/YojimboSecurity/skyplatform/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22)
[![code with love by YojimboSecurity](https://img.shields.io/badge/%3C%2F%3E%20with%20%E2%99%A5%20by-YojimboSecurity-ff1414.svg?style=flat-square)](https://github.com/YojimboSecurity)
![Go version](https://img.shields.io/badge/Go-v1.18-blue)

</div>

<details open="open">
<summary>Table of Contents</summary>

- [About](#about)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Author](#author)
- [Security](#security)
- [License](#license)

</details>

---

## About

Skyplatform is a CLI tool for Skytap. It uses Skytaps API to manage environments and virtual machines.

### Built With

To build Skyplatform, you must first install [Go](https://go.dev/). Then, you can build it with the following command:

```bash
make
./bin/Skyplatform
```

## Getting Started

To use Skyplatform, you must first create an account on [Skytap](https://www.skytap.com/). Once you have an account, you need to set environment variables.

Bash or Zsh:

```bash
export SKYTAP_USER=<email>
export SKYTAP_TOKEN=<TOKEN>
```

Fish:
  
```bash
set -gx SKYTAP_USER <email>
set -gx SKYTAP_TOKEN <TOKEN>
```

### Installation

You can install with the following command:

```bash
go install
```

## Usage

To list all environments, you can use the following command:

```bash
skyplatform environments
```

To display information about an environment, you can use the following command:

```bash
skyplatform environment <environment_id>
```

To display information about a virtual machine, you can use the following command:

```bash
skyplatform vm <vm_id>
```

To start a virtual machine, you can use the following command:

```bash
skyplatform vm start <vm_id>
```

To stop a virtual machine, you can use the following command:

```bash
skyplatform vm stop <vm_id>
```

To suspend a virtual machine, you can use the following command:

```bash
skyplatform vm suspend <vm_id>
```

## Contributing

First off, thanks for taking the time to contribute! Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make will benefit everybody else and are **greatly appreciated**.


Please read [our contribution guidelines](docs/CONTRIBUTING.md), and thank you for being involved!

## Author

The original setup of this repository is by [YojimboSecurity](https://github.com/YojimboSecurity).

## Security

Skyplatform follows good practices of security, but 100% security cannot be assured.
Skyplatform is provided **"as is"** without any **warranty**. Use at your own risk.

_For more information and to report security issues, please refer to our [security documentation](docs/SECURITY.md)._

## License

This project is licensed under the **MIT license**.

See [LICENSE](LICENSE) for more information.
