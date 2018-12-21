<p align="center">
<img width="200" alt="TensorTask Logo" src="https://storage.googleapis.com/tensortask-static/tensortask_transparent.png">
</p>

[![GoDoc][1]][2] [![Go Report Card][3]][4] [![Keybase Chat][5]][6] [![Cloud Build][7]][8]

[1]: https://godoc.org/github.com/tensortask/gotfpb?status.svg
[2]: https://godoc.org/github.com/tensortask/gotfpb
[3]: https://goreportcard.com/badge/github.com/tensortask/gotfpb
[4]: https://goreportcard.com/report/github.com/tensortask/gotfpb
[5]: https://img.shields.io/badge/keybase%20chat-tensortask.public-blue.svg
[6]: https://keybase.io/team/tensortask.public
[7]: https://storage.googleapis.com/tensortask-static/build/gotfpb.svg
[8]: https://github.com/sbsends/cloud-build-badge

# 📜 GoTFPB: Curated Golang TensorFlow Protocol Buffers

```diff
- #############################
- #       ALPHA PACKAGE       #
- #############################
```

## Package Layout

Each package has two sub-directories. The gen directory contains the generated Golang source code. The protos directory contains the protocol buffer definitions (proto3).

* **core**: fundamental protocol buffers used at the core of TensorFlow.
  * **framework**: basic absolutely vital tensors used in TensorFlow (graph, tensor, etcetera). 

## Imports

Every package imported by a Go program must end in /gen. 

* **core**:
  * **framework**: `github.com/tensortask/gotfpb/core/framework/gen`

## Compilation 

If you haven't already, download and install the [protobuf compiler](https://github.com/google/protobuf/releases) (protoc). Add the binary to your path. 

GOTFPB adds Golang-specific optimizations. These optimizations depend on a 3rd party fork of Google's protoc-gen code which improves the quality of code generated by the protoc tool. A useful generator script has been made to simplify compilation even further. Install third-party dependencies and then run generate in the repo root. 

```bash
# generate necessary code files
go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
go generate ./...
```
