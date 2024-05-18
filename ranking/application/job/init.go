package job

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/robfig/cron/v3"
	"newsCenter/idl/rankingGrpc"
	"newsCenter/ranking/infrastructure/pkg/redis_lock"
	"time"
)

func InitRankingJob(svc rankingGrpc.RankingServiceServer, client *redis_lock.Client) *RankingJob {
	return NewRankingJob(svc, client, time.Second*30)
}

func InitJobs(rjob *RankingJob) *cron.Cron {
	builder := NewCronJobBuilder(prometheus.SummaryOpts{
		Namespace: "热榜模型",
		Subsystem: "newsCenter",
		Name:      "cron_job",
		Help:      "定时任务执行",
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.75:  0.01,
			0.9:   0.01,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})
	expr := cron.New(cron.WithSeconds())
	_, err := expr.AddJob("@every 20s", builder.Build(rjob))
	if err != nil {
		panic(err)
	}
	return expr
}
