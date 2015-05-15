# go-discogs
A Go wrapper for the Discogs API.

# Example usage
```go
package main

import (
        "fmt"
        "log"

        "github.com/jesseward/go-discogs/discogs"
)

func main() {

        discogsClient := discogs.NewClient(nil)
        artist, _, err := discogsClient.Artist.Get(4130)

        if err != nil {
                log.Fatal("Error fetching artist, err=%v\n", err)
        }

        fmt.Printf("    Name : %s\n", artist.Name)
        fmt.Printf("     URI : %s\n", artist.URI)
        fmt.Printf("Releases : %s\n", artist.ReleasesURL)
}
```

Will output
```
    Name : Blunted Dummies
     URI : https://www.discogs.com/artist/4130-Blunted-Dummies
Releases : https://api.discogs.com/artists/4130/releases
```
