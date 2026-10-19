package di

import (
	"errors"
	"testing"
)

func TestContainer_ResolveRegistered(t *testing.T) {
	c := New()
	c.Register("answer", func(c *Container) (any, error) {
		return 42, nil
	})

	v, err := c.Resolve("answer")
	if err != nil {
		t.Fatal(err)
	}
	if v.(int) != 42 {
		t.Fatalf("ожидалось 42, получили %v", v)
	}
}

func TestContainer_ResolveUnknown(t *testing.T) {
	c := New()

	_, err := c.Resolve("missing")
	if err == nil {
		t.Fatal("ожидалась ошибка")
	}
}

func TestContainer_CachesResult(t *testing.T) {
	c := New()
	calls := 0
	c.Register("obj", func(c *Container) (any, error) {
		calls++
		return &struct{}{}, nil
	})

	v1, _ := c.Resolve("obj")
	v2, _ := c.Resolve("obj")

	if v1 != v2 {
		t.Fatal("ожидался один и тот же экземпляр")
	}
	if calls != 1 {
		t.Fatalf("фабрика вызвана %d раз, ожидался 1", calls)
	}
}

func TestContainer_PropagatesFactoryError(t *testing.T) {
	c := New()
	sentinel := errors.New("сломалось")
	c.Register("bad", func(c *Container) (any, error) {
		return nil, sentinel
	})

	_, err := c.Resolve("bad")
	if !errors.Is(err, sentinel) {
		t.Fatalf("ожидалась ошибка sentinel, получили %v", err)
	}
}

func TestContainer_DoesNotCacheOnError(t *testing.T) {
	c := New()
	calls := 0
	c.Register("flaky", func(c *Container) (any, error) {
		calls++
		if calls == 1 {
			return nil, errors.New("первый раз не вышло")
		}
		return "ok", nil
	})

	if _, err := c.Resolve("flaky"); err == nil {
		t.Fatal("ожидалась ошибка на первом вызове")
	}
	v, err := c.Resolve("flaky")
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != "ok" {
		t.Fatalf("ожидалось ok, получили %v", v)
	}
}

func TestContainer_RecursiveResolve(t *testing.T) {
	c := New()

	c.Register("a", func(c *Container) (any, error) {
		return "A", nil
	})
	c.Register("b", func(c *Container) (any, error) {
		a, err := c.Resolve("a")
		if err != nil {
			return nil, err
		}
		return a.(string) + "B", nil
	})

	v, err := c.Resolve("b")
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != "AB" {
		t.Fatalf("ожидалось AB, получили %v", v)
	}
}

func TestContainer_RegisterOverwrites(t *testing.T) {
	c := New()
	c.Register("x", func(c *Container) (any, error) {
		return "first", nil
	})
	c.Register("x", func(c *Container) (any, error) {
		return "second", nil
	})

	v, err := c.Resolve("x")
	if err != nil {
		t.Fatal(err)
	}
	if v.(string) != "second" {
		t.Fatalf("ожидалось second, получили %v", v)
	}
}

func TestContainer_MustResolvePanicsOnMissing(t *testing.T) {
	c := New()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидалась паника")
		}
	}()

	c.MustResolve("missing")
}
