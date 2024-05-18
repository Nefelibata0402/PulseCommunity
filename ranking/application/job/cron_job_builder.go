package job

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type CronJobBuilder struct {
	vector *prometheus.SummaryVec
}

func NewCronJobBuilder(opt prometheus.SummaryOpts) *CronJobBuilder {
	vector := prometheus.NewSummaryVec(opt,
		[]string{"job", "success"})
	return &CronJobBuilder{
		vector: vector}
}

func (b *CronJobBuilder) Build(job Job) cron.Job {
	name := job.Name()
	return cronJobAdapterFunc(func() {
		// 接入 tracing
		start := time.Now()
		zap.L().Info("开始运行", zap.String("name", name))
		err := job.Run()
		if err != nil {
			zap.L().Error("执行失败", zap.Error(err), zap.String("name", name))
		}
		zap.L().Info("结束运行", zap.String("name", name))
		duration := time.Since(start)
		b.vector.WithLabelValues(name, strconv.FormatBool(err == nil)).
			Observe(float64(duration.Milliseconds()))
	})
}

type cronJobAdapterFunc func()

func (c cronJobAdapterFunc) Run() {
	c()
}
