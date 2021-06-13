package mongo

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Dialer
//counterfeiter:generate . Client
//counterfeiter:generate . Database
//counterfeiter:generate . Collection
//counterfeiter:generate . Cursor
//counterfeiter:generate . SingleResult
