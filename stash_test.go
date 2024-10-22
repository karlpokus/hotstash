package main

import "testing"

func TestStashHit(t *testing.T) {
	stash := newStash()
	k := "key"
	v := []byte("value")
	v2 := []byte("value2")
	t.Run("miss", func(t *testing.T) {
		hit := stash.Hit(k, v)
		if hit {
			t.Fatal("unexpected stash hit")
		}
	})
	t.Run("hit", func(t *testing.T) {
		hit := stash.Hit(k, v)
		if !hit {
			t.Fatal("unexpected stash miss")
		}
	})
	t.Run("miss post update", func(t *testing.T) {
		hit := stash.Hit(k, v2)
		if hit {
			t.Fatal("unexpected stash hit")
		}
	})
	t.Run("hit post update", func(t *testing.T) {
		hit := stash.Hit(k, v2)
		if !hit {
			t.Fatal("unexpected stash miss")
		}
	})
}
