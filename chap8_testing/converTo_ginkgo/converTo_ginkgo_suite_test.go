package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConverToGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConverToGinkgo Suite")
}
