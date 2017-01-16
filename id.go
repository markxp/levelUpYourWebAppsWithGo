package main

import (
    "crypto/rand"
    "fmt"
)

// Source string used when generating a random identifier
const idSource ="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwyz"

// Length of idSource
// Why transform into byte??
const idSourceLen = byte(len(idSource))

func GenerateID(prefix string, length int) string {
    id:= make([]byte, length)
    rand.Read(id)
    for k, v := range id {
        id[k] = idSource[v%idSourceLen]
    }
    return fmt.Sprintf("%s_%s", prefix, string(id))
}
