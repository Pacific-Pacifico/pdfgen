#!/bin/bash
wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.4/wkhtmltox-0.12.4_linux-generic-amd64.tar.xz
sudo tar -xvf wkhtmltox-0.12.4_linux-generic-amd64.tar.xz 
sudo cp wkhtmltox/bin/wkhtmltopdf /usr/bin/