package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/madflow/skate-ipsum/generator"
)

type ServerOpt = struct {
	Address string
	Port    int
}

var serverOpt ServerOpt

func paragrapsHandler() http.Handler {
	paragraphs := generator.IpsumArray(10)
	fn := func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		err := json.NewEncoder(rw).Encode(paragraphs)
		if err != nil {
			http.Error(rw, err.Error(), 500)
			return
		}
	}
	return http.HandlerFunc(fn)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a skate ipsum api server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server is running ...")
		http.Handle("/", paragrapsHandler())

		err := http.ListenAndServe(":6969", nil)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	serverCmd.Flags().IntVarP(&serverOpt.Port, "port", "p", 6969, "The server port")
	serverCmd.Flags().StringVarP(&serverOpt.Address, "address", "a", "0.0.0.0", "The network address to bimd to")
	rootCmd.AddCommand(serverCmd)
}
