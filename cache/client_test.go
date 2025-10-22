package cache

import (
	"testing"
	"time"
)

func TestPutAndGet(t *testing.T) {
	c := NewCacheClient(2 * time.Second)

	ok := c.Put("city", "Curitiba")
	if !ok {
		t.Fatalf("expected Put to return true, but got false")
	}

	val, found := c.Get("city")
	if !found {
		t.Fatalf("expected to find the item, but did not")
	}

	if val != "Curitiba" {
		t.Errorf("expected 'Curitiba', got '%v'", val)
	}
}

func TestExpiration(t *testing.T) {
	c := NewCacheClient(500 * time.Millisecond)

	c.Put("temp", 25)
	time.Sleep(600 * time.Millisecond)
	_, found := c.Get("temp")
	if found {
		t.Error("expected item to expire, but it was still found")
	}
}

func TestDelete(t *testing.T) {
	c := NewCacheClient(2 * time.Second)

	c.Put("wind", "10 km/h")

	ok := c.Delete("wind")
	if !ok {
		t.Fatalf("expected Delete to return true, but got false")
	}

	_, found := c.Get("wind")
	if found {
		t.Error("expected item to be deleted, but it still exists")
	}
}

func TestOverwrite(t *testing.T) {
	c := NewCacheClient(2 * time.Second)

	c.Put("city", "Curitiba")
	c.Put("city", "Londrina")

	val, found := c.Get("city")
	if !found {
		t.Fatal("expected to find the item, but did not")
	}

	if val != "Londrina" {
		t.Errorf("expected 'Londrina', got '%v'", val)
	}
}
