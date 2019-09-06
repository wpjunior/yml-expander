groups:
  {{ range $_, $sample := .samples }}
  - name: service_{{ $sample.name }}_sample
    interval: {{ $sample.interval }}
    rules:
      {{ range $_, $window := $sample.buckets }}
      # absolute values
      - record: job:service_job_overview_duration_seconds_bucket:aggregated_over_{{$window}}
        expr: sum(increase(service_job_overview_duration_seconds_bucket[{{$window}}])) by (kind, le)

      - record: job:service_job_overview_total:aggregated_over_{{$window}}
        expr: sum(increase(service_job_overview_total[{{$window}}])) by (kind, status)
     {{ end }}
  {{ end }}

  - name: service_available
    rules:
      - alert: alert:service.uptime
        expr: absent(up{job=~"service-.*"})
        labels:
          severity: page
        annotations:
          message: My message
