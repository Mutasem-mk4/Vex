package main

import (
	"time"

	"github.com/user/vex/cmd"
	"github.com/user/vex/internal/mock"
)

func main() {
	// تشغيل سيرفر وهمي للتمكن من فحص الأداة بشكل تلقائي
	mock.StartMockServer(":8080")
	time.Sleep(500 * time.Millisecond)

	// نقل التحكم إلى مكتبة Cobra للاستخدام من سطر الأوامر
	cmd.Execute()
}
