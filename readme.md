# Tasū

Arbitrary code execution server using Docker in Go.

## Setup (Linux only)
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
