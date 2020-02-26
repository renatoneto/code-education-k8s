package main

import "testing"

func TestBold(t *testing.T) {
    result := bold("test")

    if result != "<b>test</b>" {
        t.Error("Incorrect bold result")
    }
}