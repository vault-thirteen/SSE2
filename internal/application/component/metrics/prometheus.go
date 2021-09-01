////////////////////////////////////////////////////////////////////////////////
//
//  Copyright © 2021 by Vault Thirteen.
//
//  All rights reserved. No part of this publication may be reproduced,
//  distributed, or transmitted in any form or by any means, including
//  photocopying, recording, or other electronic or mechanical methods,
//  without the prior written permission of the publisher, except in the case
//  of brief quotations embodied in critical reviews and certain other
//  noncommercial uses permitted by copyright law. For permission requests,
//  write to the publisher, addressed “Copyright Protected Material” at the
//  address below.
//
//  Web Site:  'https://github.com/vault-thirteen'.
//  Author:     Vault Thirteen.
//  Web Site Address is an Address in the global Computer Internet Network.
//
////////////////////////////////////////////////////////////////////////////////

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/vault-thirteen/SSE2/internal/application/config"
)

type Prometheus struct {
	registry         *prometheus.Registry
	processCollector prometheus.Collector
	goCollector      prometheus.Collector
	metrics          *Metrics
}

func NewPrometheus() (p *Prometheus, err error) {
	p = new(Prometheus)

	p.registry = prometheus.NewPedanticRegistry()

	p.metrics = new(Metrics)

	p.addIndicators()

	err = p.registerIndicators()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Prometheus) addIndicators() {
	p.processCollector = prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{})
	p.goCollector = prometheus.NewGoCollector()

	p.metrics.IncomingKafkaMessagesCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: config.MetricNamePrefixForService + "incoming_kafka_messages_count",
			Help: "...",
		},
	)

	p.metrics.IncomingConversionRequestsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: config.MetricNamePrefixForService + "incoming_conversion_requests_count",
			Help: "...",
		},
		[]string{config.MetricsLabelMimeType},
	)

	p.metrics.PendingTasksCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: config.MetricNamePrefixForService + "pending_tasks_count",
			Help: "...",
		},
	)

	p.metrics.ProcessedConversionRequestsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: config.MetricNamePrefixForService + "processed_conversion_requests_count",
			Help: "...",
		},
		[]string{config.MetricsLabelMimeType},
	)

	p.metrics.ConversionDurationByWorker = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: config.MetricNamePrefixForService + "conversion_duration_by_worker",
			Help: "...",
		},
		[]string{config.MetricsLabelMimeType},
	)

	p.metrics.ConversionDurationAsync = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: config.MetricNamePrefixForService + "conversion_duration_async",
			Help: "...",
		},
		[]string{config.MetricsLabelMimeType},
	)

	p.metrics.KafkaConsumerErrorsCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: config.MetricNamePrefixForService + "kafka_consumer_errors_count",
			Help: "...",
		},
	)

	p.metrics.KafkaConsumerGroupErrorsCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: config.MetricNamePrefixForService + "kafka_consumer_group_errors_count",
			Help: "...",
		},
	)

	p.metrics.KafkaProducerErrorsCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: config.MetricNamePrefixForService + "kafka_producer_errors_count",
			Help: "...",
		},
	)
}

func (p *Prometheus) registerIndicators() (err error) {
	err = p.registry.Register(p.processCollector)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.goCollector)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.IncomingKafkaMessagesCount)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.IncomingConversionRequestsCount)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.PendingTasksCount)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.ProcessedConversionRequestsCount)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.ConversionDurationByWorker)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.ConversionDurationAsync)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.KafkaConsumerErrorsCount)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.KafkaConsumerGroupErrorsCount)
	if err != nil {
		return err
	}

	err = p.registry.Register(p.metrics.KafkaProducerErrorsCount)
	if err != nil {
		return err
	}

	return nil
}

func (p *Prometheus) GetMetrics() *Metrics {
	return p.metrics
}

func (p *Prometheus) GetRegistry() *prometheus.Registry {
	return p.registry
}
