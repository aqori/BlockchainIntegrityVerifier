// cmd/blockchainintegrityverifier/main.go
package main

import (
"flag"
"log"
"os"

"blockchainintegrityverifier/internal/blockchainintegrityverifier"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainintegrityverifier.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
