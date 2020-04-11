#!/bin/bash
echo -n '{{define "main_css"}}' > `dirname $0`/../templates/view.html
cat `dirname $0`/../static/css/main.css >> `dirname $0`/../templates/view.html
echo '{{end}}' >> `dirname $0`/../templates/view.html

echo -n '{{define "main_js"}}' >> `dirname $0`/../templates/view.html
cat `dirname $0`/../static/js/main.js >> `dirname $0`/../templates/view.html
echo '{{end}}' >> `dirname $0`/../templates/view.html

echo -n '{{define "sample_jpg"}}data:image/jpeg;base64,' >> `dirname $0`/../templates/view.html
base64 -i `dirname $0`/../static/img/sample.jpg >> `dirname $0`/../templates/view.html
echo '{{end}}' >> `dirname $0`/../templates/view.html

echo -n '{{define "favicon"}}data:image/x-icon;base64,' >> `dirname $0`/../templates/view.html
base64 -i `dirname $0`/../static/img/sample_favicon.ico >> `dirname $0`/../templates/view.html
echo '{{end}}' >> `dirname $0`/../templates/view.html
