package websocket

import (
	"context"
	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"testing"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"encoding/json"
)

var jsonMetadata = getJSONMetadata()

func getJSONMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

// Run Once, Start Immediately
const testConfig string = `{
  "name": "websocket",
  "settings": {
		"url": "ws://localhost:4050/ws"
  },
  "handlers": [
    {
      "flowURI": "",
      "settings": {
        "handler_setting": "xxx"
      }
    }
  ]
}`


type TestRunner struct {
}

var Test action.Runner

// Run implements action.Runner.Run
func (tr *TestRunner) Run(context context.Context, action action.Action, uri string, options interface{}) (code int, data interface{}, err error) {
	log.Infof("Ran Action (Run): %v", uri)
	return 0, nil, nil
}

func (tr *TestRunner) RunAction(ctx context.Context, act action.Action, options map[string]interface{}) (results map[string]*data.Attribute, err error) {
	log.Infof("Ran Action (RunAction): %v", act)
	return nil, nil
}

func (tr *TestRunner) Execute(ctx context.Context, act action.Action, inputs map[string]*data.Attribute) (results map[string]*data.Attribute, err error) {
	log.Infof("Ran Action (Execute): %v", act)
	return nil, nil
}


func TestTrigger(t *testing.T) {

	log.Info("Testing Trigger")
	config := trigger.Config{}
	json.Unmarshal([]byte(testConfig),&config)

	factory := &wsTriggerFactory{}

	trigger := factory.New(&config)

	trigger.Start()

	defer trigger.Stop()


}