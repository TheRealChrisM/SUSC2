# SUSC2: Scalable Uninterruptible System for Command and Control

This project aims to create a Command and Control (C2) system that is resistant to network-level blocking and control server takedowns. It accomplishes this via a distributed system of control nodes that relies on pull-based activity (i.e., nodes polling other nodes) in order to transfer data.

## How to Use

```bash
go build -o susc2 ./cmd/node
./susc2
```
