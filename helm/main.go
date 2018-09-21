package main

import (
    "fmt"
    "k8s.io/helm/pkg/tlsutil"
    "os"
    "k8s.io/helm/pkg/helm"
    helm_env "k8s.io/helm/pkg/helm/environment"
    "k8s.io/helm/pkg/kube"
)

var (
    tlsCaCertFile string // path to TLS CA certificate file
    tlsCertFile   string // path to TLS certificate file
    tlsKeyFile    string // path to TLS key file
    tlsVerify     bool   // enable TLS and verify remote certificates
    tlsEnable     bool   // enable TLS

    tlsCaCertDefault = "$HELM_HOME/ca.pem"
    tlsCertDefault   = "$HELM_HOME/cert.pem"
    tlsKeyDefault    = "$HELM_HOME/key.pem"

    tillerTunnel *kube.Tunnel
    settings     helm_env.EnvSettings
)

func newClient() helm.Interface {
    options := []helm.Option{helm.Host(settings.TillerHost), helm.ConnectTimeout(settings.TillerConnectionTimeout)}

    if tlsVerify || tlsEnable {
        if tlsCaCertFile == "" {
            tlsCaCertFile = settings.Home.TLSCaCert()
        }
        if tlsCertFile == "" {
            tlsCertFile = settings.Home.TLSCert()
        }
        if tlsKeyFile == "" {
            tlsKeyFile = settings.Home.TLSKey()
        }
        debug("Key=%q, Cert=%q, CA=%q\n", tlsKeyFile, tlsCertFile, tlsCaCertFile)
        tlsopts := tlsutil.Options{KeyFile: tlsKeyFile, CertFile: tlsCertFile, InsecureSkipVerify: true}
        if tlsVerify {
            tlsopts.CaCertFile = tlsCaCertFile
            tlsopts.InsecureSkipVerify = false
        }
        tlscfg, err := tlsutil.ClientConfig(tlsopts)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(2)
        }
        options = append(options, helm.WithTLS(tlscfg))
    }
    return helm.NewClient(options...)
}


func debug(format string, args ...interface{}) {
    if settings.Debug {
        format = fmt.Sprintf("[debug] %s\n", format)
        fmt.Printf(format, args...)
    }
}

func main () {
    fmt.Println("hello")
}