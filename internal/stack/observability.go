package stack

// DockerComposeObservability contiene la plantilla para Prometheus + Grafana
const DockerComposeObservability = `version: '3.8'

services:
  # Prometheus - Sistema de monitoreo y alertas
  prometheus:
    image: prom/prometheus:latest
    container_name: observability_prometheus
    ports:
      - "9090:9090"
    volumes:
      - ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
      - ./prometheus/data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/etc/prometheus/console_libraries'
      - '--web.console.templates=/etc/prometheus/consoles'
      - '--web.enable-lifecycle'
    networks:
      - observability-network
    restart: unless-stopped

  # Grafana - Visualización de métricas
  grafana:
    image: grafana/grafana:latest
    container_name: observability_grafana
    ports:
      - "3000:3000"
    volumes:
      - ./grafana/data:/var/lib/grafana
      - ./grafana/provisioning:/etc/grafana/provisioning
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_INSTALL_PLUGINS=
    depends_on:
      - prometheus
    networks:
      - observability-network
    restart: unless-stopped

  # Node Exporter - Exportador de métricas del sistema
  node-exporter:
    image: prom/node-exporter:latest
    container_name: observability_node_exporter
    ports:
      - "9100:9100"
    command:
      - '--path.procfs=/host/proc'
      - '--path.rootfs=/rootfs'
      - '--path.sysfs=/host/sys'
      - '--collector.filesystem.mount-points-exclude=^/(sys|proc|dev|host|etc)($$|/)'
    volumes:
      - /proc:/host/proc:ro
      - /sys:/host/sys:ro
      - /:/rootfs:ro
    networks:
      - observability-network
    restart: unless-stopped

networks:
  observability-network:
    driver: bridge
`

// PrometheusConfig contiene la configuración básica de Prometheus
const PrometheusConfig = `global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'prometheus'
    static_configs:
      - targets: ['localhost:9090']

  - job_name: 'node-exporter'
    static_configs:
      - targets: ['node-exporter:9100']

  # Añade aquí más targets según necesites
  # - job_name: 'tu-aplicacion'
  #   static_configs:
  #     - targets: ['tu-app:puerto']
`

// GrafanaDatasource contiene la configuración del datasource de Prometheus
const GrafanaDatasource = `apiVersion: 1

datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus:9090
    isDefault: true
    editable: true
`

// ReadmeObservability contiene la documentación del stack
const ReadmeObservability = `# Stack de Observabilidad (Prometheus + Grafana)

## Estructura del proyecto

- **prometheus/**: Configuración y datos de Prometheus
- **grafana/**: Datos y configuración de Grafana

## Servicios incluidos

- **Prometheus**: Puerto 9090 - Sistema de monitoreo y alertas
- **Grafana**: Puerto 3000 - Visualización de métricas
- **Node Exporter**: Puerto 9100 - Métricas del sistema host

## Credenciales por defecto

### Grafana
- URL: http://localhost:3000
- Usuario: admin
- Password: admin

## Comandos útiles

### Iniciar el stack
` + "```bash" + `
docker-compose up -d
` + "```" + `

### Detener el stack
` + "```bash" + `
docker-compose down
` + "```" + `

### Ver logs
` + "```bash" + `
docker-compose logs -f
` + "```" + `

### Ver logs de un servicio específico
` + "```bash" + `
docker-compose logs -f prometheus
docker-compose logs -f grafana
` + "```" + `

## URLs de acceso

- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000
- Node Exporter: http://localhost:9100/metrics

## Configuración de Grafana

1. Accede a http://localhost:3000
2. Login con admin/admin
3. El datasource de Prometheus ya está configurado automáticamente
4. Importa dashboards desde https://grafana.com/grafana/dashboards/
   - Dashboard recomendado para Node Exporter: 1860

## Añadir métricas de tu aplicación

Edita ` + "`prometheus/prometheus.yml`" + ` y añade tu aplicación:

` + "```yaml" + `
scrape_configs:
  - job_name: 'mi-aplicacion'
    static_configs:
      - targets: ['host.docker.internal:puerto']
` + "```" + `

Reinicia Prometheus:
` + "```bash" + `
docker-compose restart prometheus
` + "```" + `

## Notas

- Los datos de Prometheus persisten en ` + "`prometheus/data/`" + `
- Los datos de Grafana persisten en ` + "`grafana/data/`" + `
- Node Exporter exporta métricas del host donde corre Docker
`

// GitignoreObservability contiene los archivos a ignorar
const GitignoreObservability = `prometheus/data/
grafana/data/
*.log
`

// createObservability creates an observability stack with Prometheus and Grafana
func createObservability() error {
	// Prompt for auto-start
	autoStart := PromptAutoStart()

	config := StackConfig{
		Name:        "Observability",
		Description: "Observability stack with Prometheus, Grafana and Node Exporter",
		ProjectDir:  "observability-stack",
		AutoStart:   autoStart,
		Ports: map[string]string{
			"Prometheus":    "9090",
			"Grafana":       "3000",
			"Node Exporter": "9100",
		},
		Dirs: []string{
			"prometheus",
			"prometheus/data",
			"grafana/data",
			"grafana/provisioning",
			"grafana/provisioning/datasources",
		},
		Files: map[string]string{
			"docker-compose.yml":                              DockerComposeObservability,
			"prometheus/prometheus.yml":                       PrometheusConfig,
			"grafana/provisioning/datasources/prometheus.yml": GrafanaDatasource,
			"README.md":  ReadmeObservability,
			".gitignore": GitignoreObservability,
		},
	}

	return GenerateStack(config)
}
