package store

import (
    "testing"
)

func TestCheckPalindrome(t *testing.T) {
    var testString = []Message{Message{Content: "qlik"}, Message{Content: "ahha"}}
    var results = make([]bool, 2)

    for i, str := range testString {
        results[i] = str.checkPalindrome()
    }

    if results[0] != false && results[1] != true {
        t.Fail()
    }
}
