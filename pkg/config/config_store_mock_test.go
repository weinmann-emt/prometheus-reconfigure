package config_test

type MockedStore struct {
}

func (MockedStore) Content() (string, error) {
	yamlString := `
global:
  scrape_interval: 60s
  scrape_timeout: 10s
  evaluation_interval: 60s

scrape_configs:
  - job_name: ci
    scheme: https
    tls_config:
      insecure_skip_verify: true
      ca_file: '/etc/prometheus/certs/Weinmann-Root-CA.crt'
    static_configs:
    - targets: [
    'emt-ci-lin-01.emt.local:9100',
    'emt-ci-lin-02.emt.local:9100',
      ]
  - job_name: helper
    scheme: https
    tls_config:
      insecure_skip_verify: true
      ca_file: '/etc/prometheus/certs/Weinmann-Root-CA.crt'
    static_configs:
    - targets: [
    'emt-ci-mon-01.emt.local:9100',

      ]
`
	return yamlString, nil
}

func (MockedStore) Update(string) error {
	return nil
}
