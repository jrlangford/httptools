package jsonresponder


import (
    "testing"
)

func TestResponsePushDetails(t *testing.T) {
    s := ResponseStatus{}
    s.PushDetails("DummyMessagge", "DummyCode")

    // Check the response body is what we expect.
    //expMessage := "DummyMessage"
    //expCode := "DummyCode"
    expLen := 1

    actualLen := len(s.Details)
    if len(s.Details) != 1 {
        t.Errorf("\nExp Length: %#v\nGot Length: %#v", expLen, actualLen)
    }
}
