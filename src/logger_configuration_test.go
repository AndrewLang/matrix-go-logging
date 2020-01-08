package logging

import (
	
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializeJsonConfiguration(t *testing.T) {
	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.DebugStyle = LogLevelStyle{"245", "24", ""}
	configuration.InfoStyle = LogLevelStyle{"85", "", "1"}
	configuration.WarnStyle = LogLevelStyle{"226", "124", "4"}
	configuration.ErrorStyle = LogLevelStyle{"166", "232", "1,4"}
	configuration.FatalStyle = LogLevelStyle{"196", "11", "7"}

	json := configuration.ToJSON()

		
	logger := NewConsoleLogger("Test_Env").Configure(configuration)
	logger.Info("Json logger configuration").
		Info(json)
}

func TestDeserializeJsonConfiguration(t *testing.T) {
	content := `
	{
			"LayoutNames": [
					"Time",
					"Level",
					"Name",
					"Indent",
					"Message"
			],
			"fileName": "",
			"fileSize": 2097152,
			"minLevel": 0,
			"useColor": true,
			"debugStyle": {
					"foreground": "245",
					"background": "24",
					"styles": ""
			},
			"infoStyle": {
					"foreground": "56",
					"background": "234",
					"styles": "1"
			},
			"warnStyle": {
					"foreground": "226",
					"background": "124",
					"styles": "4"
			},
			"errorStyle": {
					"foreground": "166",
					"background": "232",
					"styles": "1,4"
			},
			"fatalStyle": {
					"foreground": "196",
					"background": "11",
					"styles": "7"
			}
	}`

	configuration := NewLoggerConfiguration([]string{Time, Level, Name, Indent, Message})
	configuration.FromJSON(content)

	assert.Equal(t, "", configuration.FileName, "")
	assert.Equal(t, 0, configuration.MinLevel, "")
	assert.Equal(t, int64(2097152), configuration.FileSize, "")
	assert.Equal(t, true, configuration.UseColor, "")
	assert.Equal(t, "Time", configuration.LayoutNames[0], "")
	assert.Equal(t, "Level", configuration.LayoutNames[1], "")
	assert.Equal(t, "Name", configuration.LayoutNames[2], "")
	assert.Equal(t, "Indent", configuration.LayoutNames[3], "")
	assert.Equal(t, "Message", configuration.LayoutNames[4], "")

	assert.Equal(t, "245", configuration.DebugStyle.Foreground, "")
	assert.Equal(t, "24", configuration.DebugStyle.Background, "")
	assert.Equal(t, "", configuration.DebugStyle.Styles, "")

	assert.Equal(t, "56", configuration.InfoStyle.Foreground, "")
	assert.Equal(t, "234", configuration.InfoStyle.Background, "")
	assert.Equal(t, "1", configuration.InfoStyle.Styles, "")

	assert.Equal(t, "226", configuration.WarnStyle.Foreground, "")
	assert.Equal(t, "124", configuration.WarnStyle.Background, "")
	assert.Equal(t, "4", configuration.WarnStyle.Styles, "")

	assert.Equal(t, "166", configuration.ErrorStyle.Foreground, "")
	assert.Equal(t, "232", configuration.ErrorStyle.Background, "")
	assert.Equal(t, "1,4", configuration.ErrorStyle.Styles, "")

	assert.Equal(t, "196", configuration.FatalStyle.Foreground, "")
	assert.Equal(t, "11", configuration.FatalStyle.Background, "")
	assert.Equal(t, "7", configuration.FatalStyle.Styles, "")
}
