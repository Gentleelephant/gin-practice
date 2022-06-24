package initialize

import (
	"fmt"
	"gin-practice/src/global"
	"github.com/hashicorp/consul/api"
)

func RegisterToConsul(host, name, id string, port int, tags []string) error {
	defaultConfig := api.DefaultConfig()
	h := global.GlobalConfig.ConsulConfig.Host
	p := global.GlobalConfig.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	agentRegistration := new(api.AgentServiceRegistration)
	agentRegistration.ID = id
	agentRegistration.Name = name
	agentRegistration.Port = port
	agentRegistration.Tags = tags
	healthCheck := fmt.Sprintf("http://%s:%d/v1/health_check", host, port)
	check := &api.AgentServiceCheck{
		HTTP:                           healthCheck,
		Timeout:                        "3s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "5s",
	}
	agentRegistration.Check = check
	err = client.Agent().ServiceRegister(agentRegistration)
	return err
}

func GetServiceList() error {
	defaultConfig := api.DefaultConfig()
	h := global.GlobalConfig.ConsulConfig.Host
	p := global.GlobalConfig.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	services, err := client.Agent().Services()
	if err != nil {
		return err
	}
	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("-----------------")
	}
	return nil
}

func FilterService(filter string) error {
	defaultConfig := api.DefaultConfig()
	h := global.GlobalConfig.ConsulConfig.Host
	p := global.GlobalConfig.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		return err
	}
	// Service==account_web
	services, err := client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return err
	}
	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("-----------------")
	}
	return nil
}
