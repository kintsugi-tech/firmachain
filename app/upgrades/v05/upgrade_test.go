package v05_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/firmachain/firmachain/v05/app/apptesting"
	v05 "github.com/firmachain/firmachain/v05/app/upgrades/v05"
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (s *UpgradeTestSuite) SetupTest() {
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

// Ensures the test does not error out.
func (s *UpgradeTestSuite) TestUpgrade() {
	s.Setup()
	preUpgradeChecks(s)

	upgradeHeight := int64(5)
	s.ConfirmUpgradeSucceeded(v05.UpgradeName, upgradeHeight)

	postUpgradeChecks(s)
}

func preUpgradeChecks(_ *UpgradeTestSuite) {
}

func postUpgradeChecks(_ *UpgradeTestSuite) {
}
