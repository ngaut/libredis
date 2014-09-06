package client

import (
	"testing"
)

func TestPFAdd(t *testing.T) {
	r.Del("hll")
	n, err := r.PFAdd("hll", "a", "b")
	if err != nil {
		t.Error(err.Error())
	}
	if n != 1 {
		t.Fail()
	}
}

func TestPFCount(t *testing.T) {
	r.Del(encodeKeyWithTag("hll"), encodeKeyWithTag("hll2"))
	r.PFAdd(encodeKeyWithTag("hll"), "1", "2")
	r.PFAdd(encodeKeyWithTag("hll2"), "a", "1")
	n, err := r.PFCount(encodeKeyWithTag("hll"))
	if err != nil {
		t.Error(err.Error())
	}
	if n != 2 {
		t.Fail()
	}
	n, _ = r.PFCount(encodeKeyWithTag("hll"), encodeKeyWithTag("hll2"))
	if n != 3 {
		t.Fail()
	}
}

func TestPFMerge(t *testing.T) {
	r.Del(encodeKeyWithTag("hll"), encodeKeyWithTag("hll2"))
	r.PFAdd(encodeKeyWithTag("hll"), "a")
	r.PFAdd(encodeKeyWithTag("hll2"), "1")
	if err := r.PFMerge(encodeKeyWithTag("hll3"), encodeKeyWithTag("hll"), encodeKeyWithTag("hll2")); err != nil {
		t.Error(err.Error())
	}
}
