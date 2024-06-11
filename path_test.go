package beaut

import "testing"

func TestAbsolutePath(t *testing.T) {
	type testCase struct {
		Path       string
		expectedOk bool
	}

	testCases := []testCase{
		{Path: "a/b/c", expectedOk: false},
		{Path: "a/b/c/", expectedOk: false},
		{Path: "/a/b/c/d", expectedOk: true},
		{Path: "/a/b/c/d/", expectedOk: true},
	}

	for _, tc := range testCases {
		_, err := NewAbsolutePath(tc.Path)
		ok := err == nil
		if tc.expectedOk != ok {
			t.Errorf("expected %t, got %t for path %s", tc.expectedOk, ok, tc.Path)
		}
	}
}

func TestRelativePath(t *testing.T) {
	type testCase struct {
		Path       string
		expectedOk bool
	}

	testCases := []testCase{
		{Path: "a/b/c", expectedOk: true},
		{Path: "a/b/c/", expectedOk: true},
		{Path: "/a/b/c/d", expectedOk: false},
		{Path: "/a/b/c/d/", expectedOk: false},
	}

	for _, tc := range testCases {
		_, err := NewRelativePath(tc.Path)
		ok := err == nil
		if tc.expectedOk != ok {
			t.Errorf("expected %t, got %t for path %s", tc.expectedOk, ok, tc.Path)
		}
	}
}
