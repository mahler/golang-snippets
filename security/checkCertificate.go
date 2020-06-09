package main

import (
    "bytes"
    "crypto/tls"
    "crypto/x509" // https://golang.org/pkg/crypto/x509/
    "encoding/pem"
    "fmt"
    "os"
)

func GetCertificatesPEM(address string) (string, error) {
    conn, err := tls.Dial("tcp", address, &tls.Config{
        InsecureSkipVerify: true,
    })
    if err != nil {
        return "", err
    }
    defer conn.Close()
    var b bytes.Buffer
    for _, cert := range conn.ConnectionState().PeerCertificates {
        err := pem.Encode(&b, &pem.Block{
            Type: "CERTIFICATE",
            Bytes: cert.Raw,
        })
        if err != nil {
            return "", err
        }
    }
    return b.String(), nil
}

func main() {
    domainArg := os.Args[1]

    certs, err := GetCertificatesPEM(domainArg + ":443")
    if err != nil {
        fmt.Print("err %v\n", err.Error())
    }

    block, _ := pem.Decode([]byte(certs))
    cert, err := x509.ParseCertificate(block.Bytes)
    if err != nil {
        panic("failed to parse certificate: " + err.Error())
    }

//fmt.Printf("len %v : %v\n", len(certs), certs)
fmt.Println(cert.Subject)
fmt.Println(cert.Issuer)
fmt.Println(cert.NotAfter)

}
