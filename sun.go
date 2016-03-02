// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	l4g "github.com/alecthomas/log4go"
	"github.com/mattermost/platform/api"
	"github.com/mattermost/platform/manualtesting"
	"github.com/mattermost/platform/model"
	"github.com/mattermost/platform/utils"
	"github.com/mattermost/platform/web"

	// Plugins
	_ "github.com/mattermost/platform/model/gitlab"

	// Enterprise Deps
	_ "github.com/go-ldap/ldap"
)

func main() {

	utils.LoadConfig(flagConfigFile)

	if flagRunCmds {
		utils.ConfigureCmdLineLog()
	}

	pwd, _ := os.Getwd()
	l4g.Info(utils.T("mattermost.current_version"), model.CurrentVersion, model.BuildNumber, model.BuildDate, model.BuildHash)
	l4g.Info(utils.T("mattermost.entreprise_enabled"), model.BuildEnterpriseReady)
	l4g.Info(utils.T("mattermost.working_dir"), pwd)
	l4g.Info(utils.T("mattermost.config_file"), utils.FindConfigFile(flagConfigFile))

	api.NewServer()
	api.InitApi()
	web.InitWeb()

	if model.BuildEnterpriseReady == "true" {
		loadLicense()
	}

	if flagRunCmds {
		runCmds()
	} else {
		api.StartServer()

		// If we allow testing then listen for manual testing URL hits
		if utils.Cfg.ServiceSettings.EnableTesting {
			manualtesting.InitManualTesting()
		}

		setDiagnosticId()
		runSecurityAndDiagnosticsJobAndForget()

		// wait for kill signal before attempting to gracefully shutdown
		// the running service
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c

		api.StopServer()
	}
}
