#!/bin/bash
BASE_DIR=`dirname $0`/../
TARGET_FILE=${BASE_DIR}templates/view.html

echo -n '{{define "main.css"}}' > ${TARGET_FILE}
cat ${BASE_DIR}static/css/main.css >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}

echo -n '{{define "main.js"}}' >> ${TARGET_FILE}
cat ${BASE_DIR}static/js/main.js >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}

echo -n '{{define "favicon"}}data:image/x-icon;base64,' >> ${TARGET_FILE}
base64 -i ${BASE_DIR}static/img/sample_favicon.ico >> ${TARGET_FILE}
echo '{{end}}' >> ${TARGET_FILE}

IMAGES="${BASE_DIR}static/img/*"
FILEARY=()
for FILEPATH in ${IMAGES}; do
  if [ -f ${FILEPATH} ] ; then
    FILEARY+=("${FILEPATH}")
  fi
done

for i in ${FILEARY[@]}; do
  FILENAME=`basename ${i}`
  if [ ${FILENAME##*.} == jpg ] ; then
    echo -n '{{define "'${FILENAME}'"}}data:image/jpeg;base64,' >> ${TARGET_FILE}
    base64 -i ${i} >> ${TARGET_FILE}
    echo '{{end}}' >> ${TARGET_FILE}
  elif [ ${FILENAME##*.} == png ] ; then
    echo -n '{{define "'${FILENAME}'"}}data:image/png;base64,' >> ${TARGET_FILE}
    base64 -i ${i} >> ${TARGET_FILE}
    echo '{{end}}' >> ${TARGET_FILE}
  fi
done
