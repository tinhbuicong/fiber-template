##### Setup
```bash
docker compose -f deploy/docker-compose.logging.yml up -d
```

##### Hệ thống logs
```
App
 ├── Metrics → Prometheus → Alertmanager
 └── Logs    → Promtail → Loki
                     ↓
                  Grafana
```