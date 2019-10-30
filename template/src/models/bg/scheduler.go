package bg

import (
	"fmt"
	"{{name}}/src/config"
	"{{name}}/src/models/wechat"

	"github.com/jasonlvhit/gocron"
)

// DailyJob will handle daily task
func DailyJob() {
	if config.All["wechat.enable"] == "true" {
		gocron.Every(1).Minute().Do(refreshGlobalAccessToken)
	}

	_, time := gocron.NextRun()
	fmt.Println("################## next bg task will run @", time)

	// start and hold on
	<-gocron.Start()
}

func refreshGlobalAccessToken() {
	fmt.Println("################### triggered task - refresh wechat global access token ####################")

	wechat.GetOrRefreshGlobalAccessToken()
}
