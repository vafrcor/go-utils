# `GetImageConfig` Unit Test Design

## Scope

Add unit coverage for `GetImageConfig` in `common/filesystem/image_test.go`. Production code and unrelated worktree changes remain untouched.

## Test structure

Use Go's standard `testing` package and files created inside `t.TempDir()`. Generate a small PNG in the test with `image.NewRGBA` and `png.Encode`, avoiding committed binary fixtures and external dependencies.

Cover three behaviors:

1. A valid PNG returns a non-nil configuration with the encoded width, height, and RGBA color model.
2. A missing path returns a nil configuration and an error containing `can't check open file`.
3. A readable non-image file returns a nil configuration and an error containing `can't decode image`.

Use subtests so failures identify the behavior under test. Fail setup immediately with `t.Fatal`; validate returned values with ordinary `testing` comparisons.

## Verification

Run the package test for `common/filesystem`. If the repository's currently deleted module files prevent that command, report the exact limitation and verify formatting with `gofmt`.
