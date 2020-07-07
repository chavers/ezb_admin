//     This file is part of ezBastion.
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

package setup

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/chavers/ezb_admin/models"
	log "github.com/sirupsen/logrus"
)

var (
	confFile string
)

func init() {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex)
	confFile = path.Join(exPath, "conf/config.json")
}

//CheckConfig try to found config json file and parse it
func CheckConfig() (conf models.Configuration, err error) {
	raw, err := ioutil.ReadFile(confFile)
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		return conf, err
	}
	log.Debug("json config found and loaded.")
	return conf, nil

}

// func Setup(isIntSess bool) error {

// 	_fqdn := fqdn.Get()
// 	quiet := true
// 	hostname, _ := os.Hostname()
// 	conf, err := CheckConfig()
// 	if err != nil {
// 		quiet = false
// 		conf.ServiceFullName = "Easy Bastion admin console"
// 		conf.ServiceName = "ezb_adm"
// 		conf.Logger.LogLevel = "warning"
// 		conf.Logger.MaxSize = 10
// 		conf.Logger.MaxBackups = 5
// 		conf.Logger.MaxAge = 180
// 		conf.CaCert = "cert/ca.crt"
// 		conf.PrivateKey = "cert/ezb_adm.key"
// 		conf.PublicCert = "cert/ezb_adm.crt"
// 		conf.Listen = "0.0.0.0:8080"
// 		conf.EzbDB = "https://localhost:5501"
// 		conf.EzbSRV = "https://localhost:5505"
// 	}

// 	_, fica := os.Stat(path.Join(exPath, conf.CaCert))
// 	_, fipriv := os.Stat(path.Join(exPath, conf.PrivateKey))
// 	_, fipub := os.Stat(path.Join(exPath, conf.PublicCert))
// 	if quiet == false {
// 		fmt.Print("\n\n")
// 		fmt.Println("***********")
// 		fmt.Println("*** PKI ***")
// 		fmt.Println("***********")
// 		fmt.Println("ezBastion nodes use elliptic curve digital signature algorithm ")
// 		fmt.Println("(ECDSA) to communicate.")
// 		fmt.Println("We need ezb_pki address and port, to request certificat pair.")
// 		fmt.Println("ex: 10.20.1.2:6000 pki.domain.local:6000")

// 		for {
// 			p := setupmanager.AskForValue("ezb_pki", conf.EzbPki, `^[a-zA-Z0-9-\.]+:[0-9]{4,5}$`)
// 			c := setupmanager.AskForConfirmation(fmt.Sprintf("pki address (%s) ok?", p))
// 			if c {
// 				conn, err := net.Dial("tcp", p)
// 				if err != nil {
// 					fmt.Printf("## Failed to connect to %s ##\n", p)
// 				} else {
// 					conn.Close()
// 					conf.EzbPki = p
// 					break
// 				}
// 			}
// 		}

// 		fmt.Print("\n\n")
// 		fmt.Println("Certificat Subject Alternative Name.")
// 		fmt.Printf("\nBy default using: <%s, %s> as SAN. Add more ?\n", _fqdn, hostname)
// 		for {
// 			tmp := conf.SAN

// 			san := setupmanager.AskForValue("SAN (comma separated list)", strings.Join(conf.SAN, ","), `(?m)^[[:ascii:]]*,?$`)

// 			t := strings.Replace(san, " ", "", -1)
// 			tmp = strings.Split(t, ",")
// 			c := setupmanager.AskForConfirmation(fmt.Sprintf("SAN list %s ok?", tmp))
// 			if c {
// 				conf.SAN = tmp
// 				break
// 			}
// 		}
// 	}
// 	if os.IsNotExist(fica) || os.IsNotExist(fipriv) || os.IsNotExist(fipub) {
// 		keyFile := path.Join(exPath, conf.PrivateKey)
// 		certFile := path.Join(exPath, conf.PublicCert)
// 		caFile := path.Join(exPath, conf.CaCert)

// 	}
// 	if quiet == false {
// 		c, _ := json.Marshal(conf)
// 		ioutil.WriteFile(confFile, c, 0600)
// 		log.Println(confFile, " saved.")
// 	}

// 	return nil
// }
