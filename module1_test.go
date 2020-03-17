package alog

import "testing"

// 01-01 (Task 01, Test 01)
func TestMessageChannelModule1(t *testing.T) {
	alog := New(nil)
	if alog.msgCh == nil {
		t.Fatal("msgCh field not initialized. Should have type 'chan string' but it is currently nil")
	}
}

// 02-01
func TestErrorChannelModule1(t *testing.T) {
	alog := New(nil)
	if alog.errorCh == nil {
		t.Fatal("errorCh field not initialized. Should have type 'chan string' but it is currently nil")
	}
}