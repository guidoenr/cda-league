#!/bin/bash

API_TOKEN="d8ac4de15667c7c45378be669af450f4cf55b2e9"


convertImage(){
  curl \
      -F 'image=@manucho.jpg' https://cartoonify.net/cartoonize
}

convertImage