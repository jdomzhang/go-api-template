package bg

import (
	"fmt"
	"{{name}}/src/config"
	"{{name}}/src/models/biz"

	"github.com/jasonlvhit/gocron"
)

// DailyJob will handle daily task
func DailyJob() {
	if config.All["wechat.enable"] == "true" {
		gocron.Every(1).Day().At("3:00").Do(deleteExpiredFormID)
		gocron.Every(1).Do(wechat.GetOrRefreshGlobalAccessToken)
	}

	_, time := gocron.NextRun()
	fmt.Println("################## next bg task will run @", time)

	// start and hold on
	<-gocron.Start()
}

func deleteExpiredFormID() {
	fmt.Println("################### triggered daily task - delete expired form id ####################")

	biz.DeleteExpiredFormID()
}
