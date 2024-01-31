package test

import (
	"testing"
	"github.com/onsi/gomega"
	"github.com/Level69Magikrap/se_lab_test01/entity"
)
func NameTest(t *testing.T){
	g := gomega.NewGomegaWithT(t)
	t.Run("need name"),func(t *testing.T){
		g := entity.User(){
			Name
		}
	}
}