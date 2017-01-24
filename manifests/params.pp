class prometheus::params {
  $blackbox_download_url = 'https://github.com/jetstack-experimental/blackbox_exporter/releases/download/poc-proxy/'
  $blackbox_dest_dir = '/usr/local/sbin'
  $systemd_path = '/etc/systemd/system'
  $addon_dir = '/etc/kubernetes/addons'
  $helper_dir = '/usr/local/sbin'
  $node_exporter_image = 'prom/node-exporter'
  $node_exporter_version = 'v0.13.0'
  $node_exporter_port = 9190
  $prometheus_namespace = 'monitoring'
  $prometheus_image = 'quay.io/prometheus/prometheus'
  $prometheus_version = 'v1.4.1'
  $prometheus_port = 9090
  $prometheus_storage_local_retention = '6h'
  $prometheus_storage_local_memchunks = 500000
  $prometheus_use_module_config = true
  $prometheus_use_module_rules = true
  $prometheus_install_state_metrics = true
  $prometheus_install_node_exporter = true
}
