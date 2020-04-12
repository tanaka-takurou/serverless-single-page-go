#!/bin/bash
BASE_DIR=`dirname $0`/../
TARGET_FILE=${BASE_DIR}templates/view.html

echo -n '{{define "main_css"}}' > ${TARGET_FILE}
cat ${BASE_DIR}static/css/main.css >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}

echo -n '{{define "main_js"}}' >> ${TARGET_FILE}
cat ${BASE_DIR}static/js/main.js >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}

echo -n '{{define "sample_jpg"}}data:image/jpeg;base64,' >> ${TARGET_FILE}
base64 -i ${BASE_DIR}static/img/sample.jpg >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}

echo -n '{{define "favicon"}}data:image/x-icon;base64,' >> ${TARGET_FILE}
base64 -i ${BASE_DIR}static/img/sample_favicon.ico >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}
