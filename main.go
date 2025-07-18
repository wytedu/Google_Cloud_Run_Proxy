package main
import (
        "io"
        "net"
        "os"
)
func main() {
        listenAddr := ":" + os.Getenv("PORT")
        targetAddr := os.Getenv("165.232.175.163") + ":45639"
        ln, err := net.Listen("tcp", listenAddr)
        if err != nil {
                return
        }
        for {
                conn, err := ln.Accept()
                if err != nil {
                        continue
                }
                go handleConnection(conn, targetAddr)
        }
}
func handleConnection(src net.Conn, targetAddr string) {
        dst, err := net.Dial("tcp", targetAddr)
        if err != nil {
                src.Close()
                return
        }
        go io.Copy(dst, src)
        go io.Copy(src, dst)
}
