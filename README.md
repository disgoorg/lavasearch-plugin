# lavasearch-plugin

[DisGoLink](https://github.com/disgoorg/disgolink) plugin for [LavaSearch](https://github.com/topi314/LavaSearch) for advanced [Lavalink](https://github.com/lavalink-devs/Lavalink) search.

## Installation

```shell
$ go get github.com/disgoorg/disgo
```

## Usage

```go
var client disgolink.Client
result, err := lavasearch.LoadSearch(client.BestNode().Rest(), "dzsearch:test", []lavasearch.SearchType{lavasearch.SearchTypeAlbum, lavasearch.SearchTypePlaylist})
if err != nil {
    // handle error
}

result.Albums // []lavasearch.Album
result.Playlists // []lavasearch.Playlist
```

## Troubleshooting

For help feel free to open an issue or reach out on [Discord](https://discord.gg/TewhTfDpvW)

## Contributing

Contributions are welcomed but for bigger changes we recommend first reaching out via [Discord](https://discord.gg/TewhTfDpvW) or create an issue to discuss your problems, intentions and ideas.

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/disgoorg/disgo/blob/master/LICENSE). See LICENSE for more information.
