package job

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"x-ui/logger"
	"x-ui/web/service"
)

type SendTrafficDataJob struct {
	inboundService service.InboundService
}

func NewSendTrafficDataJob() *SendTrafficDataJob {
	return new(SendTrafficDataJob)
}

func (j *SendTrafficDataJob) Run() {
	logger.Info("Send info to Matreshka back...")

	apiUrl := "https://tuatara.space/api/v1/server/vray-statistic"

	inboundList, err := j.inboundService.GetAllInbounds()

	if err != nil {
		logger.Warning("Error on getting inbounds")
		return
	}

	dataForSend, err := json.Marshal(inboundList)

	if err != nil {
		logger.Warning("Error on build Json")
		return
	}

	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer([]byte(dataForSend)))

	if err != nil {
		logger.Warning("Error on sent request to Matreshka")
		return
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("Status: ", response.Status)

	// clean up memory after execution
	defer response.Body.Close()
}
