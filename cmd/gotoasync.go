/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
    "async/avro"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/swaggest/go-asyncapi/reflector/asyncapi-2.0.0"
	"github.com/swaggest/go-asyncapi/spec-2.0.0"


)

// gotoasyncCmd represents the gotoasync command
var gotoasyncCmd = &cobra.Command{
	Use:   "gotoasync",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gotoasync called")
		//broker_url := "https://pkc-0wg55.us-central1.gcp.devel.cpdev.cloud:443"
		url := "https://psrc-8jz90.us-west-2.aws.devel.cpdev.cloud/subjects"
		req, _ := http.NewRequest("GET", url, nil)
		req.SetBasicAuth("PAU2KX7Z6LIROIQK", "8Qm4emI/yvsg0G/lYxrtNuQmSVUOsrSQ/MYVE8bvy83DgC6AR3le3q5pbMs89PyP")

		resp, _ := http.DefaultClient.Do(req)
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		abc := string(body)
		abc = abc[1 : len(abc)-1]
		res1 := strings.Split(abc, ",")
		reflector := asyncapi.Reflector{
			Schema: &spec.AsyncAPI{
				Servers: map[string]spec.Server{
					"production": {
						URL:             "https://pkc-0wg55.us-central1.gcp.devel.cpdev.cloud:443",
						Description:     "Kafka Production instance.",
						ProtocolVersion: "7.2.0",
						Protocol:        "Kafka",
					},
				},
				Info: spec.Info{
					Version: "2.0.0", // required
					Title:   "API Document for Confluent Cluster",
				},
			},
		}
		mustNotFail := func(err error) {
			if err != nil {
				panic(err.Error())
			}
		}
		//schema := `{"type":"record","name":"Msg1","namespace":"com.mycorp.mynamespace","doc":"Sample schema to help you get started.","fields":[{"name":"orderId","type":"int","doc":"The id of the order."},{"name":"orderTime","type":"int","doc":"Timestamp of the order."},{"name":"orderAddress","type":"string","doc":"The address of the order."}]}`

        for i:=0;i<2;i++ {
            mustNotFail(reflector.AddChannel(asyncapi.ChannelInfo{
            			Name: res1[i],
            			Publish: &asyncapi.MessageSample{
            				MessageSample: new(avro.Msg),
            			},
            		}))


        }
        yaml, err := reflector.Schema.MarshalYAML()
                    		mustNotFail(err)
		fmt.Println(string(yaml))
		mustNotFail(ioutil.WriteFile("async.yaml", yaml, 0644))

	},
}

func init() {
	rootCmd.AddCommand(gotoasyncCmd)

}
