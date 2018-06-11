package application

import (
	"github.com/ThomasFPienaar/collaborative/modules"
	"github.com/ThomasFPienaar/collaborative/application/httplayer"
)

func InitApplication() {
	modules.LoadModules()
}

func Start() {
	httplayer.SetupHttpServer()

}