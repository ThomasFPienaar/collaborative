package application

import (
	"github.com/thomas/collaborative/modules"
	"github.com/thomas/collaborative/application/httplayer"
)

func InitApplication() {
	modules.LoadModules()
}

func Start() {
	httplayer.SetupHttpServer()

}