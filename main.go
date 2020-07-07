//go:generate goversioninfo

// This file is part of ezBastion.
//     ezBastion is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Affero General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//     ezBastion is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Affero General Public License for more details.
//     You should have received a copy of the GNU Affero General Public License
//     along with ezBastion.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/urfave/cli/v2"
)

var (
	exPath string
	// conf   models.Configuration
)

const (
	name        = "ezb_admin"
	description = "ezBastion web admin console."
	displayName = "ezBastion web admin console."
)

func init() {
	ex, _ := os.Executable()
	exPath = filepath.Dir(ex)
}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {

	router := gin.Default()
	router.Static("/", exPath+string(os.PathSeparator)+"docs")
	router.Run(":8080")
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {

	svcConfig := &service.Config{
		Name:             name,
		DisplayName:      displayName,
		Description:      description,
		WorkingDirectory: exPath,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if service.Interactive() {

		app := cli.NewApp()
		app.Name = "ezb_admin"
		app.Version = "0.2.1"
		app.Usage = "ezBastion web admin console."

		app.Commands = []*cli.Command{
			// {
			// 	Name:  "init",
			// 	Usage: "Genarate config file.",
			// 	Action: func(c *cli.Context) error {
			// 		return setup.Setup()
			// 	},
			// },
			{
				Name:  "debug",
				Usage: "Start ezb_adm in console.",
				Action: func(c *cli.Context) error {
					return s.Run()
				},
			},
			{
				Name:  "install",
				Usage: "Add ezb_adm deamon windows service.",
				Action: func(c *cli.Context) error {
					err = s.Install()
					if err == nil {
						logger.Info("service installed")
					}
					return err
				},
			},
			{
				Name:  "remove",
				Usage: "Remove ezb_adm deamon windows service.",
				Action: func(c *cli.Context) error {
					err = s.Uninstall()
					if err == nil {
						logger.Info("service uninstalled")
					}
					return err
				},
			}, {
				Name:  "start",
				Usage: "Start ezb_adm deamon windows service.",
				Action: func(c *cli.Context) error {
					err = s.Start()
					if err == nil {
						logger.Info("service started")
					}
					return err
				},
			}, {
				Name:  "stop",
				Usage: "Stop ezb_adm deamon windows service.",
				Action: func(c *cli.Context) error {
					err = s.Stop()
					if err == nil {
						logger.Info("service stop")
					}
					return err
				},
			},
		}
		cli.AppHelpTemplate = fmt.Sprintf(`

	███████╗███████╗██████╗  █████╗ ███████╗████████╗██╗ ██████╗ ███╗   ██╗
	██╔════╝╚══███╔╝██╔══██╗██╔══██╗██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║
	█████╗    ███╔╝ ██████╔╝███████║███████╗   ██║   ██║██║   ██║██╔██╗ ██║
	██╔══╝   ███╔╝  ██╔══██╗██╔══██║╚════██║   ██║   ██║██║   ██║██║╚██╗██║
	███████╗███████╗██████╔╝██║  ██║███████║   ██║   ██║╚██████╔╝██║ ╚████║
	╚══════╝╚══════╝╚═════╝ ╚═╝  ╚═╝╚══════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝
																		   
							 █████╗ ██████╗ ███╗   ███╗                    
							██╔══██╗██╔══██╗████╗ ████║                    
							███████║██║  ██║██╔████╔██║                    
							██╔══██║██║  ██║██║╚██╔╝██║                    
							██║  ██║██████╔╝██║ ╚═╝ ██║                    
							╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝                    

%s
INFO:
		http://www.ezbastion.com		
		support@ezbastion.com
		`, cli.AppHelpTemplate)
		err = app.Run(os.Args)
		if err != nil {
			logger.Error(err)
		}
	} else {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
}
