package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/spf13/cobra"
)

var gpioCmd = &cobra.Command{Use:"gpio",Short:"GPIO Pins"}

var gpioReadCmd = &cobra.Command{
	Use:"read [pin]",Short:"GPIO lesen",Args:cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command,args []string) error {
		resp,err:=http.Get(apiURL()+"/api/v1/gpio/"+args[0])
		if err!=nil{return err}
		defer resp.Body.Close()
		var d map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&d)
		sym:="🔴 LOW"; if d["value"].(float64)==1{sym="🟢 HIGH"}
		fmt.Printf("Pin %-8s → %s\n",args[0],sym)
		return nil
	},
}

var gpioWriteCmd = &cobra.Command{
	Use:"write [pin] [0|1]",Short:"GPIO setzen",Args:cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command,args []string) error {
		val,err:=strconv.Atoi(args[1])
		if err!=nil||(val!=0&&val!=1){return fmt.Errorf("Wert muss 0 oder 1 sein")}
		body,_:=json.Marshal(map[string]int{"value":val})
		resp,err:=http.Post(apiURL()+"/api/v1/gpio/"+args[0],"application/json",bytes.NewReader(body))
		if err!=nil{return err}
		resp.Body.Close()
		sym:="🔴 LOW"; if val==1{sym="🟢 HIGH"}
		fmt.Printf("✅ Pin %s → %s\n",args[0],sym)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gpioCmd)
	gpioCmd.AddCommand(gpioReadCmd)
	gpioCmd.AddCommand(gpioWriteCmd)
}
