package common

func Must[T any](v T, e error) T {
  if e != nil {
    panic(e)
  }
  return v
}

func Must2[T any, P any](v1 T, v2 P, e error) (T, P) {
  if e != nil {
    panic(e)
  }
  return v1, v2
}