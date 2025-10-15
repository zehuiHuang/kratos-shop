package data_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/data"
)

var _ = Describe("User", func() {
	var ro biz.UserRepo
	var uD *biz.User
	BeforeEach(func() {
		// 这里的 Db 是 data_suite_test.go 文件里面定义的
		ro = data.NewUserRepo(Db, nil)
		// 这里你可以引入外部组装好的数据
		uD = &biz.User{
			ID:       1,
			Mobile:   "13803881388",
			Password: "admin123456",
			NickName: "aliliin",
			Role:     1,
			Birthday: 693629981,
		}
	})

	// 设置 It 块来添加单个规格
	It("CreateUser", func() {
		u, err := ro.CreateUser(ctx, uD)
		Ω(err).ShouldNot(HaveOccurred())
		// 组装的数据 mobile 为 13803881388
		Ω(u.Mobile).Should(Equal("13803881388")) // 手机号应该为创建的时候写入的手机号
	})

})
