# chroma-hub

Graphics hub built using [Golang][go].
Stores the graphics templates for instances of [Chroma Viz][chroma-viz] and [Chroma Engine][chroma-engine].
This ensures data synchronization between instances and one point of updating for changes.

**[ARCHIVED]** Chroma Hub now lives as a submodule in the [Chroma Viz][chroma-viz] repo.

## Installation

- Download source code and run

```
go run main.go -port [9000] -file [file]
```

## Disclaimer

This is a personal project, not an application intended for production.

[go]: https://github.com/golang/go
[chroma-engine]: https://github.com/jchilds0/chroma-engine
[chroma-viz]: https://github.com/jchilds0/chroma-viz
