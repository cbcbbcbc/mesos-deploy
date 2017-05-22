#!/usr/bin/env bash


BASE_DIR=$(cd `dirname $0` && pwd -P)



curl -L -O https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-5.4.0-amd64.deb
dpkg -i filebeat-5.4.0-amd64.deb && rm -rf filebeat-5.4.0-amd64.deb

cp ${BASE_DIR}/filebeat/config/filebeat.yml /etc/filebeat/filebeat.yml
sed -i -e "s#master#${MASTER_IP}#g" /etc/filebeat/filebeat.yml
systemctl restart filebeat
systemctl enable filebeat
systemctl status filebeat


curl -L -O https://artifacts.elastic.co/downloads/beats/metricbeat/metricbeat-5.4.0-amd64.deb
dpkg -i metricbeat-5.4.0-amd64.deb  && rm -rf metricbeat-5.4.0-amd64.deb


cp ${BASE_DIR}/metricbeat/config/metricbeat.yml /etc/metricbeat/metricbeat.yml
sed -i -e "s#master#${MASTER_IP}#g" /etc/metricbeat/metricbeat.yml
systemctl restart metricbeat
systemctl enable metricbeat
systemctl status metricbeat

/usr/share/filebeat/scripts/import_dashboards -es http://${MASTER_IP}:9200
/usr/share/metricbeat/scripts/import_dashboards -es http://${MASTER_IP}:9200
