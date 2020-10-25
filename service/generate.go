package service

//go:generate goderive .
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . DependencyIterator
//counterfeiter:generate . Service
//counterfeiter:generate . Listener
