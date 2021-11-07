package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Marc"},
			[]string{"Marc"},
		},
		{
			"Struct with two string fields",
			struct {
				Name     string
				Nickname string
			}{"Marc", "Hamilton"},
			[]string{"Marc", "Hamilton"},
		},
		{
			"Struct with a non-string field",
			struct {
				Name string
				Age  int
			}{"Marc", 49},
			[]string{"Marc"},
		},
		{
			"Struct with nested struct",
			Person{
				"Marc",
				Profile{49, "Hamilton"},
			},
			[]string{"Marc", "Hamilton"},
		},
		{
			"Pointers to things",
			&Person{
				"Marc",
				Profile{49, "Hamilton"},
			},
			[]string{"Marc", "Hamilton"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{32, "Regina"},
			},
			[]string{"London", "Regina"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{32, "Regina"},
			},
			[]string{"London", "Regina"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("expected %q, got %q", test.Expected, got)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		m := map[string]string{
			"one": "Foo",
			"two": "Bar",
		}

		var got []string
		walk(m, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Foo")
		assertContains(t, got, "Bar")
	})

	t.Run("Channels", func(t *testing.T) {
		c := make(chan Profile)

		go func() {
			c <- Profile{33, "London"}
			c <- Profile{32, "Regina"}
			close(c)
		}()

		var got []string
		expected := []string{"London", "Regina"}

		walk(c, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %q, got %q", expected, got)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, got []string, expected string) {
	t.Helper()
	contains := false

	for _, x := range got {
		if x == expected {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %q to contain %s", got, expected)
	}
}
