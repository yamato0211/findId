package findId_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/yamato0211/findId"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, findId.Analyzer, "a")
}
