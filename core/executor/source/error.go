package source

import (
	"github.com/baojiweicn/Surge/util/errors"
)

var (
	// SourceNotExistsError : is the error for package manager not founc
	SourceNotExistsError = errors.NewError("environment {{language}} package manager not exists")
	// PackageNotInstalledError : is the error for the package not installed
	PackageNotInstalledError = errors.NewError("{{package}} not installed")

	// PackageVersionNotMatchError : the package installed but the version not match
	PackageVersionNotMatchError = errors.NewError("{{package}} current version {{current}} but want {{want}}")
)