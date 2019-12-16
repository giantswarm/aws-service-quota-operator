package project

var (
	bundleVersion = "0.0.1"
	description   = "The aws-service-quota-operator does something."
	gitSHA        = "n/a"
	name          = "aws-service-quota-operator"
	source        = "https://github.com/giantswarm/aws-service-quota-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
