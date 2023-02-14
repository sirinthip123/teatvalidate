package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`
	Url  string `gorm:"uniqueIndex" valid:"url"`
}


func TestVideoValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check data is valid", func(t *testing.T) {
		u := Video{
			Name: "Chatree Nernplub",
			Url:  "https://www.youtube.com/",
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		//g.Expect(err)
	})
}


func TestTest(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check name not blank", func(t *testing.T) {
		u := Video{
			Name: "", //ผิด
			Url:  "https://www.youtube.com/",
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name cannot be blank"))
	})
}

func Test(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check Url is valid", func(t *testing.T) {
		u := Video{
			Name: "Chatree",
			Url:  "://www.youtubes.com/", //ผิด 
		}
		ok, err := govalidator.ValidateStruct(u)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Url: ://www.youtubes.com/ does not validate as url"))
		
	})


}