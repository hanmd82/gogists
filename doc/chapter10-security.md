## Security

**Generate a Self-Signed TLS Certificate**
- use the `crypto/tls` package in Go’s standard library
    ```bash
    go run /usr/local/Cellar/go/1.15.2/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
    # 2020/10/22 09:17:01 wrote cert.pem
    # 2020/10/22 09:17:01 wrote key.pem
    ```
