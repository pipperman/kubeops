package server

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type ServerOption func(ops *kubeOps)

func WithProjectOption(manager ProjectManagerServer) ServerOption {
	return func(ops *kubeOps) {
		ops.projectManager = manager
	}
}

func WithRunnerOption(manager RunnerManagerServer) ServerOption {
	return func(ops *kubeOps) {
		ops.runnerManager = manager
	}
}

func WithPoolOption(taskQueueSize, workerSize int) ServerOption {
	return func(ops *kubeOps) {
		ops.pool = NewPool(ops.ctx, taskQueueSize, workerSize)
	}
}

func WithTaskCacheOption(defaultExpiration, cleanupInterval time.Duration) ServerOption {
	return func(ops *kubeOps) {
		ops.taskCache = cache.New(defaultExpiration, cleanupInterval)
	}
}

func WithDataCacheOption(defaultExpiration, cleanupInterval time.Duration) ServerOption {
	return func(ops *kubeOps) {
		ops.dataCache = cache.New(defaultExpiration, cleanupInterval)
	}
}

func WithInventoryCacheOption(defaultExpiration, cleanupInterval time.Duration) ServerOption {
	return func(ops *kubeOps) {
		ops.inventoryCache = cache.New(defaultExpiration, cleanupInterval)
	}
}