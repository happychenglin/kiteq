package server

import (
	"kiteq/remoting"
	"kiteq/stat"
	"time"
)

type KiteQConfig struct {
	flowstat          *stat.FlowStat
	rc                *remoting.RemotingConfig
	server            string
	zkhost            string
	deliverTimeout    time.Duration //投递超时时间
	maxDeliverWorkers int           //最大执行实际那
	recoverPeriod     time.Duration //recover的周期
	topics            []string      //可以处理的topics列表
	db                string        //持久层配置
}

func NewKiteQConfig(flowstat *stat.FlowStat, server, zkhost string, deliverTimeout time.Duration, maxDeliverWorkers int,
	recoverPeriod time.Duration,
	topics []string,
	db string,
	rc *remoting.RemotingConfig) KiteQConfig {
	rc.FlowStat = flowstat.RemotingFlow
	return KiteQConfig{
		flowstat:          flowstat,
		rc:                rc,
		server:            server,
		zkhost:            zkhost,
		deliverTimeout:    deliverTimeout,
		maxDeliverWorkers: maxDeliverWorkers,
		recoverPeriod:     recoverPeriod,
		topics:            topics,
		db:                db}
}
