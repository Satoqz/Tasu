<br />
<p align="center">
  <a href="https://github.com/github_username/repo_name">
    <!-- <img src="images/logo.png" alt="Logo" width="80" height="80"> !-->
  </a>

  <h1 align="center">Tasū</h1>

  <p align="center">
    Arbitrary code execution server using Docker written in Go.
    <br />
    <strong>
		Inspired by
		<a href="https://github.com/1Computer1/Myriad#readme">Myriad</a>
		and
		<a href="https://github.com/iCrawl/Myrias#readme">Myrias</a>
		»
	</strong>
    <br />
    <br />
  </p>
</p>


## Table of Contents

* [Setup](#setup)
  * [Prepacked binaries/executables](#prepacked-binaries/executables)
  * [Config](#config)
  * [gVisor](#gvisor)
* [Routes](#routes)
* [Command line arguments](#command-line-arguments)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)

## Setup

### Prepacked binaries/executables
Check the releases tab to donwload the prepacked binaries for windows/linux. They also come with the languages folder included to get you started immediately.

### Config
You'll need to create a config.json to configure Tasu as following:

`config.json`:
```json
{
	"languages": [
		"cpp",
		"javascript",
		"python"
	],
	"cleanupInterval": "10m",
	"ram": 128,
	"swap": 128
}
```

- ram and swap in `mb`
- cleanupInterval determines in which time interval all containers are either restarted or started again if crashed, arguments such as `30s`, `10m`, `1h` work here
- for a list of supported languages check the `/languages` folder, feel free to add your own languages or change the current behavior

### gVisor/runsc setup
**This is relevant for linux only**
You have to [install](https://gvisor.dev/docs/user_guide/docker/) [gVisor](https://github.com/google/gvisor) as a runtime for docker to provide an additional isolation boundary between the containers and the host kernel.

```sh
(
    set -e 
    wget https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc
    wget https://storage.googleapis.com/gvisor/releases/nightly/latest/runsc.sha512
    sha512sum -c runsc.sha512
    sudo mv runsc /usr/local/bin
    sudo chown root:root /usr/local/bin/runsc
    sudo chmod 0755 /usr/local/bin/runsc
)
```

`/etc/docker/daemon.json`:
```json
{
    "runtimes": {
        "runsc": {
            "path": "/usr/local/bin/runsc",
            "runtimeArgs": [
                "--network=none",
                "--overlay"
            ]
        },
        "runsc-kvm": {
            "path": "/usr/local/bin/runsc",
            "runtimeArgs": [
                "--platform=kvm",
                "--network=none",
                "--overlay"
            ]
        }
    }
}
```
You may have to create this file if it does not exist.

### Command line arguments
You can pass `--buildContainers` or `-bc` to build all containers for the appropriate languages in your config on startup.

## Routes

### POST `/eval`
Evaluate code.<br>
Example request json body:<br>
```json
{
	"code": "print('hello world')",
	"language": "python"
}
```
Example response #1 (success)
```json
{
	"output": "hello world\n"
}
```
Example response #2 (container restarting)
```json
{
	"error": "Currently waiting for container restart"
}
```
Example response #2 (container unavailable due to another reason)
```json
{
	"error": "Container currently unavailable"
}
```
- Other error responses can be expected if an invalid language or invalid json is received
### GET `/languages`
List of languages in config.<br>
Example response:
```json
["go", "typescript", "rust"]
```

### GET `/status`
List of all tasu docker containers and their status.<br>
Example response:
```json
[
	"tasu_go": {
		"alive": true,
		"language": "go"
	},
	"tasu_cpp": {
		"alive": false,
		"language": "cpp"
	}
]
```

### GET `/kill`
Kills all containers and prevents further restarting.<br>
Responds with the names of the killed containers.<br>
Example response:
```json
["tasu_go", "tasu_javascript", "tasu_rust"]
```
## Contributing
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Contact
Project Link: [https://github.com/Satoqz/Tasu](https://github.com/Satoqz/Tasu)
