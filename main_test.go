package main

import "testing"

func TestSimpleFactory(t *testing.T) {
    f := SimpleFactory("http://localhost")

    if f.Host != "http://localhost" {
        t.Errorf("Host incorrect, got %s, want: %s", f.Host, "http://localhost")
    }
}
