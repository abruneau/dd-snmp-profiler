[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <!-- <a href="https://github.com/abruneau/dd-snmp-profiler">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <h3 align="center">dd-snmp-profiler</h3>

  <p align="center">
    SNMP profiler generator for Datadog
    <br />
    dd-snmp-profiler is a CLI to generate SNMP profiles for Datadog
    <br />
    <a href="https://github.com/abruneau/dd-snmp-profiler"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/abruneau/dd-snmp-profiler">View Demo</a>
    ·
    <a href="https://github.com/abruneau/dd-snmp-profiler/issues">Report Bug</a>
    ·
    <a href="https://github.com/abruneau/dd-snmp-profiler/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
<!-- * [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation) -->
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
<!-- * [Acknowledgements](#acknowledgements) -->



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

<!-- Here's a blank template to get started:
**To avoid retyping too much info. Do a search and replace with your text editor for the following:**
`abruneau`, `dd-snmp-profiler`, `BruneauAntonin`, `antonin.bruneau@gmail.com` -->


### Built With

* [cobra](https://github.com/spf13/cobra)
* [gosmi](github.com/sleepinggenius2/gosmi)
* [yaml](https://github.com/go-yaml/yaml)



<!-- GETTING STARTED -->
<!-- ## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
```sh
npm install npm@latest -g
```

### Installation

1. Clone the repo
```sh
git clone https://github.com/abruneau/dd-snmp-profiler.git
```
2. Install NPM packages
```sh
npm install
```
 -->


<!-- USAGE EXAMPLES -->
## Usage

`help` command
```sh
dd-snmp-profiler help
dd-snmp-profiler is a CLI to generate SNMP profiles for Datadog

Usage:
  dd-snmp-profiler [command]

Available Commands:
  help        Help about any command
  mib         Generate profile from MIB file

Flags:
  -d, --debug           ennable debug mode
  -h, --help            help for dd-snmp-profiler
  -o, --output string   name of the output file

Use "dd-snmp-profiler [command] --help" for more information about a command.
```

`mib` command
```sh
dd-snmp-profiler mib --help
mib (dd-snmp-profiler mib) will generate an new snmp profile
        for Datadog agent based on the provided MIB file

Usage:
  dd-snmp-profiler mib [flags]

Flags:
  -h, --help            help for mib
  -m, --module string   Module to load
  -p, --path string     Path to add

Global Flags:
  -d, --debug           ennable debug mode
  -o, --output string   name of the output file
```

example:
```sh
dd-snmp-profiler mib -d -p /usr/share/snmp/mibs -m AIRPORT-BASESTATION-3-MIB -o test.yaml  
Loaded module AIRPORT-BASESTATION-3-MIB
Wrote AIRPORT-BASESTATION-3-MIB profile to test.yaml
```

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/abruneau/dd-snmp-profiler/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Antonin Bruneau - [@BruneauAntonin](https://twitter.com/BruneauAntonin) - antonin.bruneau@gmail.com

Project Link: [https://github.com/abruneau/dd-snmp-profiler](https://github.com/abruneau/dd-snmp-profiler)



<!-- ACKNOWLEDGEMENTS
## Acknowledgements

* []()
* []()
* []() -->





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/abruneau/dd-snmp-profiler.svg?style=flat-square
[contributors-url]: https://github.com/abruneau/dd-snmp-profiler/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/abruneau/dd-snmp-profiler.svg?style=flat-square
[forks-url]: https://github.com/abruneau/dd-snmp-profiler/network/members
[stars-shield]: https://img.shields.io/github/stars/abruneau/dd-snmp-profiler.svg?style=flat-square
[stars-url]: https://github.com/abruneau/dd-snmp-profiler/stargazers
[issues-shield]: https://img.shields.io/github/issues/abruneau/dd-snmp-profiler.svg?style=flat-square
[issues-url]: https://github.com/abruneau/dd-snmp-profiler/issues
[license-shield]: https://img.shields.io/github/license/abruneau/dd-snmp-profiler.svg?style=flat-square
[license-url]: https://github.com/abruneau/dd-snmp-profiler/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/antoninbruneau
[product-screenshot]: images/screenshot.png